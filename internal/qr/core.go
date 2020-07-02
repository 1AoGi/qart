package qr

import (
	"github.com/nfnt/resize"
	"github.com/tautcony/qart/models/qr"
	"github.com/tautcony/qart/models/request"
	"image"
	"image/color"
	_ "image/png"
)

func Draw(op *request.Operation, i image.Image) (*qr.Image, error) {
	target := makeTarget(i, 17+4*op.Version+op.Size)

	qrImage := &qr.Image{
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

	if err := qrImage.Encode(); err != nil {
		return nil, err
	}
	return qrImage, nil
}

func makeTarget(i image.Image, max int) [][]int {
	b := i.Bounds()
	dx, dy := max, max
	if b.Dx() > b.Dy() {
		dy = b.Dy() * dx / b.Dx()
	} else {
		dx = b.Dx() * dy / b.Dy()
	}
	thumbnail := resize.Resize(uint(dx), uint(dy), i, resize.Bilinear)

	b = thumbnail.Bounds()
	dx, dy = b.Dx(), b.Dy()
	target := make([][]int, dy)
	arr := make([]int, dx*dy)
	for y := 0; y < dy; y++ {
		target[y], arr = arr[:dx], arr[dx:]
		row := target[y]
		for x := 0; x < dx; x++ {
			p := thumbnail.At(x, y)
			_, _, _, a := p.RGBA()
			luminance := color.Gray16Model.Convert(p).(color.Gray16)
			if a == 0 {
				row[x] = -1
			} else {
				row[x] = int(luminance.Y >> 8)
			}
		}
	}
	return target
}
