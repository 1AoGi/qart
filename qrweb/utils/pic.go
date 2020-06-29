package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"rsc.io/qr/coding"
)

// todo: add caption
func MakeImage(caption, font string, pt, size, border, scale int, f func(x, y int) uint32) *image.RGBA {
	d := (size + 2*border) * scale
	csize := 0
	if caption != "" {
		if pt == 0 {
			pt = 11
		}
		csize = pt * 2
	}
	c := image.NewRGBA(image.Rect(0, 0, d, d+csize))

	// white
	u := &image.Uniform{C: color.White}
	draw.Draw(c, c.Bounds(), u, image.ZP, draw.Src)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			r := image.Rect((x+border)*scale, (y+border)*scale, (x+border+1)*scale, (y+border+1)*scale)
			rgba := f(x, y)
			u.C = color.RGBA{byte(rgba >> 24), byte(rgba >> 16), byte(rgba >> 8), byte(rgba)}
			draw.Draw(c, r, u, image.ZP, draw.Src)
		}
	}

	/*if csize != 0 {
		if font == "" {
			font = "data/luxisr.ttf"
		}
		ctxt := fs.NewContext(req)
		dat, _, err := ctxt.Read(font)
		if err != nil {
			panic(err)
		}
		tfont, err := freetype.ParseFont(dat)
		if err != nil {
			panic(err)
		}
		ft := freetype.NewContext()
		ft.SetDst(c)
		ft.SetDPI(100)
		ft.SetFont(tfont)
		ft.SetFontSize(float64(pt))
		ft.SetSrc(image.NewUniform(color.Black))
		ft.SetClip(image.Rect(0, 0, 0, 0))
		wid, err := ft.DrawString(caption, freetype.Pt(0, 0))
		if err != nil {
			panic(err)
		}
		p := freetype.Pt(d, d+3*pt/2)
		p.X -= wid.X
		p.X /= 2
		ft.SetClip(c.Bounds())
		ft.DrawString(caption, p)
	}
	*/

	return c
}

func makeFrame(font string, pt, vers, l, scale, dots int) image.Image {
	lev := coding.Level(l)
	p, err := coding.NewPlan(coding.Version(vers), lev, 0)
	if err != nil {
		panic(err)
	}

	nd := p.DataBytes / p.Blocks
	nc := p.CheckBytes / p.Blocks
	extra := p.DataBytes - nd*p.Blocks

	cap := fmt.Sprintf("QR v%d, %s", vers, lev)
	if dots > 0 {
		cap = fmt.Sprintf("QR v%d order, from bottom right", vers)
	}
	m := MakeImage(cap, font, pt, len(p.Pixel), 0, scale, func(x, y int) uint32 {
		pix := p.Pixel[y][x]
		switch pix.Role() {
		case coding.Data:
			if dots > 0 {
				return 0xffffffff
			}
			off := int(pix.Offset() / 8)
			nd := nd
			var i int
			for i = 0; i < p.Blocks; i++ {
				if i == extra {
					nd++
				}
				if off < nd {
					break
				}
				off -= nd
			}
			return blockColors[i%len(blockColors)]
		case coding.Check:
			if dots > 0 {
				return 0xffffffff
			}
			i := (int(pix.Offset()/8) - p.DataBytes) / nc
			return dark(blockColors[i%len(blockColors)])
		}
		if pix&coding.Black != 0 {
			return 0x000000ff
		}
		return 0xffffffff
	})

	if dots > 0 {
		b := m.Bounds()
		for y := 0; y <= len(p.Pixel); y++ {
			for x := 0; x < b.Dx(); x++ {
				m.SetRGBA(x, y*scale-(y/len(p.Pixel)), color.RGBA{127, 127, 127, 255})
			}
		}
		for x := 0; x <= len(p.Pixel); x++ {
			for y := 0; y < b.Dx(); y++ {
				m.SetRGBA(x*scale-(x/len(p.Pixel)), y, color.RGBA{127, 127, 127, 255})
			}
		}
		order := make([]image.Point, (p.DataBytes+p.CheckBytes)*8+1)
		for y, row := range p.Pixel {
			for x, pix := range row {
				if r := pix.Role(); r != coding.Data && r != coding.Check {
					continue
				}
				//	draw.Draw(m, m.Bounds().Add(image.Pt(x*scale, y*scale)), dot, image.ZP, draw.Over)
				order[pix.Offset()] = image.Point{x*scale + scale/2, y*scale + scale/2}
			}
		}

		for mode := 0; mode < 2; mode++ {
			for i, p := range order {
				q := order[i+1]
				if q.X == 0 {
					break
				}
				line(m, p, q, mode)
			}
		}
	}
	return m
}

func line(m *image.RGBA, p, q image.Point, mode int) {
	x := 0
	y := 0
	dx := q.X - p.X
	dy := q.Y - p.Y
	xsign := +1
	ysign := +1
	if dx < 0 {
		xsign = -1
		dx = -dx
	}
	if dy < 0 {
		ysign = -1
		dy = -dy
	}
	pt := func() {
		switch mode {
		case 0:
			for dx := -2; dx <= 2; dx++ {
				for dy := -2; dy <= 2; dy++ {
					if dy*dx <= -4 || dy*dx >= 4 {
						continue
					}
					m.SetRGBA(p.X+x*xsign+dx, p.Y+y*ysign+dy, color.RGBA{255, 192, 192, 255})
				}
			}

		case 1:
			m.SetRGBA(p.X+x*xsign, p.Y+y*ysign, color.RGBA{128, 0, 0, 255})
		}
	}
	if dx > dy {
		for x < dx || y < dy {
			pt()
			x++
			if float64(x)*float64(dy)/float64(dx)-float64(y) > 0.5 {
				y++
			}
		}
	} else {
		for x < dx || y < dy {
			pt()
			y++
			if float64(y)*float64(dx)/float64(dy)-float64(x) > 0.5 {
				x++
			}
		}
	}
	pt()
}

func PngEncode(c image.Image) []byte {
	var b bytes.Buffer
	err := png.Encode(&b, c)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

var blockColors = []uint32{
	0x7777ffff,
	0xffff77ff,
	0xff7777ff,
	0x77ffffff,
	0x1e90ffff,
	0xffffe0ff,
	0x8b6969ff,
	0x77ff77ff,
	0x9b30ffff,
	0x00bfffff,
	0x90e890ff,
	0xfff68fff,
	0xffec8bff,
	0xffa07aff,
	0xffa54fff,
	0xeee8aaff,
	0x98fb98ff,
	0xbfbfbfff,
	0x54ff9fff,
	0xffaeb9ff,
	0xb23aeeff,
	0xbbffffff,
	0x7fffd4ff,
	0xff7a7aff,
	0x00007fff,
}

func dark(x uint32) uint32 {
	r, g, b, a := byte(x>>24), byte(x>>16), byte(x>>8), byte(x)
	r = r/2 + r/4
	g = g/2 + g/4
	b = b/2 + b/4
	return uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a)
}

func clamp(x int) byte {
	if x < 0 {
		return 0
	}
	if x > 255 {
		return 255
	}
	return byte(x)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
