package qr

import (
	"bytes"
	"fmt"
	"net/http"
	"rsc.io/qr/coding"
	"rsc.io/qr/gf256"
)

type BitBlock struct {
	DataBytes  int
	CheckBytes int
	B          []byte
	M          [][]byte
	Tmp        []byte
	RS         *gf256.RSEncoder
	bdata      []byte
	cdata      []byte
}

func newBlock(nd, nc int, rs *gf256.RSEncoder, dat, cdata []byte) *BitBlock {
	b := &BitBlock{
		DataBytes:  nd,
		CheckBytes: nc,
		B:          make([]byte, nd+nc),
		Tmp:        make([]byte, nc),
		RS:         rs,
		bdata:      dat,
		cdata:      cdata,
	}
	copy(b.B, dat)
	rs.ECC(b.B[:nd], b.B[nd:])
	b.check()
	if !bytes.Equal(b.Tmp, cdata) {
		panic("cdata")
	}

	b.M = make([][]byte, nd*8)
	for i := range b.M {
		row := make([]byte, nd+nc)
		b.M[i] = row
		for j := range row {
			row[j] = 0
		}
		row[i/8] = 1 << (7 - uint(i%8))
		rs.ECC(row[:nd], row[nd:])
	}
	return b
}

func (b *BitBlock) check() {
	b.RS.ECC(b.B[:b.DataBytes], b.Tmp)
	if !bytes.Equal(b.B[b.DataBytes:], b.Tmp) {
		fmt.Printf("ecc mismatch\n%x\n%x\n", b.B[b.DataBytes:], b.Tmp)
		panic("mismatch")
	}
}

func (b *BitBlock) reset(bi uint, bval byte) {
	if (b.B[bi/8]>>(7-bi&7))&1 == bval {
		// already has desired bit
		return
	}
	// rows that have already been set
	m := b.M[len(b.M):cap(b.M)]
	for _, row := range m {
		if row[bi/8]&(1<<(7-bi&7)) != 0 {
			// Found it.
			for j, v := range row {
				b.B[j] ^= v
			}
			return
		}
	}
	panic("reset of unset bit")
}

func (b *BitBlock) canSet(bi uint, bval byte) bool {
	found := false
	m := b.M
	for j, row := range m {
		if row[bi/8]&(1<<(7-bi&7)) == 0 {
			continue
		}
		if !found {
			found = true
			if j != 0 {
				m[0], m[j] = m[j], m[0]
			}
			continue
		}
		for k := range row {
			row[k] ^= m[0][k]
		}
	}
	if !found {
		return false
	}

	targ := m[0]

	// Subtract from saved-away rows too.
	for _, row := range m[len(m):cap(m)] {
		if row[bi/8]&(1<<(7-bi&7)) == 0 {
			continue
		}
		for k := range row {
			row[k] ^= targ[k]
		}
	}

	// Found a row with bit #bi == 1 and cut that bit from all the others.
	// Apply to data and remove from m.
	if (b.B[bi/8]>>(7-bi&7))&1 != bval {
		for j, v := range targ {
			b.B[j] ^= v
		}
	}
	b.check()
	n := len(m) - 1
	m[0], m[n] = m[n], m[0]
	b.M = m[:n]

	for _, row := range b.M {
		if row[bi/8]&(1<<(7-bi&7)) != 0 {
			panic("did not reduce")
		}
	}

	return true
}

func (b *BitBlock) copyOut() {
	b.check()
	copy(b.bdata, b.B[:b.DataBytes])
	copy(b.cdata, b.B[b.DataBytes:])
}

func showtable(w http.ResponseWriter, b *BitBlock, gray func(int) bool) {
	nd := b.DataBytes
	nc := b.CheckBytes

	fmt.Fprintf(w, "<table class='matrix' cellspacing=0 cellpadding=0 border=0>\n")
	line := func() {
		fmt.Fprintf(w, "<tr height=1 bgcolor='#bbbbbb'><td colspan=%d>\n", (nd+nc)*8)
	}
	line()
	dorow := func(row []byte) {
		fmt.Fprintf(w, "<tr>\n")
		for i := 0; i < (nd+nc)*8; i++ {
			fmt.Fprintf(w, "<td")
			v := row[i/8] >> uint(7-i&7) & 1
			if gray(i) {
				fmt.Fprintf(w, " class='gray'")
			}
			fmt.Fprintf(w, ">")
			if v == 1 {
				fmt.Fprintf(w, "1")
			}
		}
		line()
	}

	m := b.M[len(b.M):cap(b.M)]
	for i := len(m) - 1; i >= 0; i-- {
		dorow(m[i])
	}
	m = b.M
	for _, row := range b.M {
		dorow(row)
	}

	fmt.Fprintf(w, "</table>\n")
}

func BitsTable(w http.ResponseWriter, req *http.Request) {
	nd := 2
	nc := 2
	fmt.Fprintf(w, `<html>
		<style type='text/css'>
		.matrix {
			font-family: sans-serif;
			font-size: 0.8em;
		}
		table.matrix {
			padding-left: 1em;
			padding-right: 1em;
			padding-top: 1em;
			padding-bottom: 1em;
		}
		.matrix td {
			padding-left: 0.3em;
			padding-right: 0.3em;
			border-left: 2px solid white;
			border-right: 2px solid white;
			text-align: center;
			color: #aaa;
		}
		.matrix td.gray {
			color: black;
			background-color: #ddd;
		}
		</style>
	`)
	rs := gf256.NewRSEncoder(coding.Field, nc)
	dat := make([]byte, nd+nc)
	b := newBlock(nd, nc, rs, dat[:nd], dat[nd:])
	for i := 0; i < nd*8; i++ {
		b.canSet(uint(i), 0)
	}
	showtable(w, b, func(i int) bool { return i < nd*8 })

	b = newBlock(nd, nc, rs, dat[:nd], dat[nd:])
	for j := 0; j < (nd+nc)*8; j += 2 {
		b.canSet(uint(j), 0)
	}
	showtable(w, b, func(i int) bool { return i%2 == 0 })

}
