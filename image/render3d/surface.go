package render3d

import (
	"fmt"
	"io"
	"math"
)

type SurfaceOption struct {
	Width  int
	Height int
	Cells  int
	Range  float64
	Scale  float64
	Zscale float64
	Angle  float64
}

var defaultSurfaceOption = &SurfaceOption{
	Width:  600,
	Height: 320,
	Cells:  100,
	Range:  30.0,
	Scale:  1.0,
	Zscale: 0.4,
	Angle:  math.Pi / 6,
}

func Surface(w io.Writer, option *SurfaceOption) {
	if option == nil {
		option = defaultSurfaceOption
	}

	width, height, cells := option.Width, option.Height, option.Cells
	xyrange, xyscale, zscale := option.Range, (float64(width)/2/option.Range)*option.Scale, float64(height)*option.Zscale
	sin, cos := math.Sin(option.Angle), math.Cos(option.Angle)
	corner := func(i, j int) (float64, float64) {
		x := xyrange * (float64(i)/float64(cells) - 0.5)
		y := xyrange * (float64(j)/float64(cells) - 0.5)

		z := f(x, y)

		sx := float64(width)/2 + (x-y)*cos*xyscale
		sy := float64(height)/2 + (x+y)*sin*xyscale - z*zscale
		return sx, sy
	}

	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			s += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	s += fmt.Sprintln("</svg>")
	w.Write([]byte(s))
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
