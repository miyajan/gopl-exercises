package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	err := png.Encode(os.Stdout, img)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		z = z - (z*z*z*z-1)/(4*z*z*z)
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}
