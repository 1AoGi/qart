package utils

import (
	"github.com/nfnt/resize"

	"image"
	_ "image/jpeg"
	"io"
)

func GetImageThumbnail(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	// Convert image to 256x256 size.
	b := img.Bounds()
	const max = 256

	// Encode to PNG.
	dx, dy := max, max
	if b.Dx() > b.Dy() {
		dy = b.Dy() * dx / b.Dx()
	} else {
		dx = b.Dx() * dy / b.Dy()
	}
	i128 := resize.Resize(uint(dx), uint(dy), img, resize.Bicubic)

	return i128, nil
}
