package render2d

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var defaultPalette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
}

const (
	whiteIndex = 0
	blackIndex = 1
	redIndex   = 2
	greenIndex = 3
	blueIndex  = 4
)

var defaultLissajousOption = &LissajousOption{
	Cycles:  5,
	Size:    100,
	Nframes: 64,
	Delay:   8,
	Res:     0.001,
}

type LissajousOption struct {
	Cycles  int
	Size    int
	Nframes int
	Delay   int
	Res     float64
}

func Lissajous(out io.Writer, palette []color.Color, option *LissajousOption) {
	if palette == nil {
		palette = defaultPalette
	}
	if option == nil {
		option = defaultLissajousOption
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: option.Nframes}
	phase, size := 0.0, option.Size
	for i := 0; i < anim.LoopCount; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(option.Cycles)*2*math.Pi; t += option.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blueIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, option.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
