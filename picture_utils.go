package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
)

// GenerateAndShow generates a 256x256 pixel base64 encoded image and
// also prints it to the command line
func GenerateAndShow(f func(int, int) [][]uint8, print bool) (string, error) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = v
			m.Pix[i+3] = v
		}
	}
	encoded, err := GenerateBase64String(m)
	if err != nil {
		return "", err
	}

	if print {
		fmt.Println("IMAGE:" + encoded)
	}
	return encoded, nil
}

// GenerateBase64String generates a Base64 representation of an image
func GenerateBase64String(m image.Image) (string, error) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, m); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
