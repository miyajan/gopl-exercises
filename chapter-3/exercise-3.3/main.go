package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, aok := corner(i+1, j)
			bx, by, bz, bok := corner(i, j)
			cx, cy, cz, cok := corner(i, j+1)
			dx, dy, dz, dok := corner(i+1, j+1)
			if aok && bok && cok && dok {
				z := (az + bz + cz + dz) / 4
				red := int((z/2 + 0.5) * 255)
				blue := int((-z/2 + 0.5) * 255)
				color := fmt.Sprintf("#%02x00%02x", red, blue)
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					"style='stroke: %s; fill: %s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
