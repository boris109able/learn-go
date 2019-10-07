package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

// green RGBA(0, 128, 0, 255)
var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0x80, 0x00, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2
)

func lissajous(out io.Writer, cycles int) {
	const (
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animations frames
		delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles) * 2.0 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //NOTE: ignoring encoding errors
}

// handler adjusts cycles according to http://localhost:8000/?cycles=20
func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	if len(r.URL.RawQuery) > 7 {
		if i, err := strconv.Atoi(r.URL.RawQuery[7:]); err == nil {
			cycles = i
		}
	}
	lissajous(w, cycles)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


