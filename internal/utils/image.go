package utils

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"github.com/tautcony/qart/internal/resize"
)

func GetImageThumbnail(r io.Reader) (bytes.Buffer, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return bytes.Buffer{}, err
	}
	// Convert image to 128x128 gray+alpha.
	b := img.Bounds()
	const max = 128
	// If it's gigantic, it's more efficient to downsample first
	// and then resize; resizing will smooth out the roughness.
	var i1 *image.RGBA
	if b.Dx() > 4*max || b.Dy() > 4*max {
		w, h := 2*max, 2*max
		if b.Dx() > b.Dy() {
			h = b.Dy() * h / b.Dx()
		} else {
			w = b.Dx() * w / b.Dy()
		}
		i1 = resize.Resample(img, b, w, h)
	} else {
		// "Resample" to same size, just to convert to RGBA.
		i1 = resize.Resample(img, b, b.Dx(), b.Dy())
	}
	b = i1.Bounds()

	// Encode to PNG.
	dx, dy := 128, 128
	if b.Dx() > b.Dy() {
		dy = b.Dy() * dx / b.Dx()
	} else {
		dx = b.Dx() * dy / b.Dy()
	}
	i128 := resize.ResizeRGBA(i1, i1.Bounds(), dx, dy)

	var buf bytes.Buffer
	if err := png.Encode(&buf, i128); err != nil {
		return bytes.Buffer{}, err
	}

	return buf, nil
}
