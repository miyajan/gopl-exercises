package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	wi := r.Form.Get("width")
	width, err := strconv.Atoi(wi)
	if err != nil {
		width = 600
	}
	h := r.Form.Get("height")
	height, err := strconv.Atoi(h)
	if err != nil {
		height = 320
	}
	c := r.Form.Get("color")

	w.Header().Set("Content-Type", "image/svg+xml")
	svg(w, float64(width), float64(height), c)
}

func svg(out io.Writer, width float64, height float64, color string) {
	if color == "" {
		color = "white"
	}

	_, err := fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int(width), int(height))
	if err != nil {
		fmt.Fprintf(os.Stderr, "svg: %v\n", err)
		return
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aok := corner(width, height, i+1, j)
			bx, by, bok := corner(width, height, i, j)
			cx, cy, cok := corner(width, height, i, j+1)
			dx, dy, dok := corner(width, height, i+1, j+1)
			if aok && bok && cok && dok {
				_, err = fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					"style='stroke: grey; fill: %s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
				if err != nil {
					fmt.Fprintf(os.Stderr, "svg: %v\n", err)
					return
				}
			}
		}
	}
	_, err = fmt.Fprint(out, "</svg>")
	if err != nil {
		fmt.Fprintf(os.Stderr, "svg: %v\n", err)
		return
	}
}

func corner(width float64, height float64, i, j int) (float64, float64, bool) {
	xyscale := width / 2 / xyrange
	zscale := height * 0.4

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
