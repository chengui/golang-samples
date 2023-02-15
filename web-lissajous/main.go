package main

import (
    "log"
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "net/http"
    "os"
    "time"
    "strconv"
)

var palette = []color.Color{
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

var (
    cycles  = 5
    size    = 100
    nframes = 64
    delay   = 8
    res     = 0.001
)

func main() {
    if false {
        rand.Seed(time.Now().UTC().UnixNano())
        lissajous(os.Stdout)
    } else {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            if err := r.ParseForm(); err != nil {
                log.Print(err)
            }
            for k, v := range r.Form {
                switch k {
                case "cycles":
                    cycles, _ = strconv.Atoi(v[0])
                case "size":
                    size, _ = strconv.Atoi(v[0])
                case "nframes":
                    nframes, _ = strconv.Atoi(v[0])
                case "delay":
                    delay, _ = strconv.Atoi(v[0])
                }
            }
            lissajous(w)
        })
        log.Fatal(http.ListenAndServe("localhost:8000", nil))
    }
}

func lissajous(out io.Writer) {
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blueIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
