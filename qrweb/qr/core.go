package qr

import (
	"bytes"
	"image"
	"path/filepath"
	"qart/models/qr"
	"qart/models/request"
	"qart/qrweb/resize"
	"qart/qrweb/utils"
)

func Draw(op *request.Operation) ([]byte, error) {
	target := makeTarget(op.Image, 17+4*op.Version+op.Size)

	img := &qr.Image{
		Name:         op.Image,
		Dx:           op.Dx,
		Dy:           op.Dy,
		URL:          op.URL,
		Version:      op.GetVersion(),
		Mask:         op.Mask,
		RandControl:  op.RandControl,
		Dither:       op.Dither,
		OnlyDataBits: op.OnlyDataBits,
		SaveControl:  op.SaveControl,
		Scale:        op.GetScale(),
		Target:       target,
		Seed:         op.GetSeed(),
		Rotation:     op.GetRotation(),
		Size:         op.Size,
	}

	if err := img.Encode(); err != nil {
		return nil, err
	}

	var dat []byte
	switch {
	case img.SaveControl:
		dat = img.Control
	default:
		dat = img.Code.PNG()
	}
	return dat, nil
}

func loadSize(name string, max int) *image.RGBA {
	data, _, err := utils.Read(getUploadPath(name + ".png"))
	if err != nil {
		panic(err)
	}
	i, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	b := i.Bounds()
	dx, dy := max, max
	if b.Dx() > b.Dy() {
		dy = b.Dy() * dx / b.Dx()
	} else {
		dx = b.Dx() * dy / b.Dy()
	}
	var irgba *image.RGBA
	switch i := i.(type) {
	case *image.RGBA:
		irgba = resize.ResizeRGBA(i, i.Bounds(), dx, dy)
	case *image.NRGBA:
		irgba = resize.ResizeNRGBA(i, i.Bounds(), dx, dy)
	}
	return irgba
}

func makeTarget(name string, max int) [][]int {
	i := loadSize(name, max)
	b := i.Bounds()
	dx, dy := b.Dx(), b.Dy()
	targ := make([][]int, dy)
	arr := make([]int, dx*dy)
	for y := 0; y < dy; y++ {
		targ[y], arr = arr[:dx], arr[dx:]
		row := targ[y]
		for x := 0; x < dx; x++ {
			p := i.Pix[y*i.Stride+4*x:]
			r, g, b, a := p[0], p[1], p[2], p[3]
			if a == 0 {
				row[x] = -1
			} else {
				row[x] = int((299*uint32(r) + 587*uint32(g) + 114*uint32(b) + 500) / 1000)
			}
		}
	}
	return targ
}

func getStoragePath(elem ...string) string {
	return filepath.Join("storage", filepath.Join(elem...))
}

func getUploadPath(name string) string {
	return getStoragePath("upload", name)
}
