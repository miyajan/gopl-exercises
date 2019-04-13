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
	dx := 1.0 / width * (xmax - xmin)
	dy := 1.0 / height * (ymax - ymin)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, superSampling(x, y, dx, dy))
		}
	}
	err := png.Encode(os.Stdout, img)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}

func superSampling(x float64, y float64, dx float64, dy float64) color.Color {
	c1 := mandelbrot(complex(x - dx, y - dy))
	c2 := mandelbrot(complex(x - dx, y + dy))
	c3 := mandelbrot(complex(x + dx, y - dy))
	c4 := mandelbrot(complex(x + dx, y + dy))
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	r3, g3, b3, a3 := c3.RGBA()
	r4, g4, b4, a4 := c4.RGBA()
	r := uint8(((r1 + r2 + r3 + r4) / 4) >> 8)
	g := uint8(((g1 + g2 + g3 + g4) / 4) >> 8)
	b := uint8(((b1 + b2 + b3 + b4) / 4) >> 8)
	a := uint8(((a1 + a2 + a3 + a4) / 4) >> 8)
	return color.RGBA{R: r, G: g, B: b, A: a}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}
