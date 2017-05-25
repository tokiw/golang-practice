package main

import (
	"image"
	"image/color"
	"math/cmplx"
)

func img(centerX, centerY, zoom float64) *image.RGBA {
	const (
		width, height = 1024, 1024
	)
	var xmin, ymin, xmax, ymax = -zoom, -zoom, +zoom, +zoom

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x+centerX, y+centerY)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := contrast * n
			b := 255 - r
			g := 0
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.Black
}
