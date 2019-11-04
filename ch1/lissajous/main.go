// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"strconv"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

type LissajousParams struct {
	Cycles  int     // number of complete x oscillator revolutions
	Res     float64 // angular resolution
	Size    int     // image canvas covers [-size..+size]
	Nframes int     // number of animation frames
	Delay   int     // delay between frames in 10ms units
}

var defaultLissajousParams = LissajousParams {
	Cycles  : 5,
	Res     : 0.001,
	Size    : 100,
	Nframes : 64,
	Delay   : 8,
}

var palette = []color.Color{
	color.White,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.Black,
}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			params := defaultLissajousParams
			if query["cycles"] != nil {
				params.Cycles, _ = strconv.Atoi(query.Get("cycles"))
			}
			if query["res"] != nil {
				params.Res, _ = strconv.ParseFloat(query.Get("res"), 64)
			}
			if query["size"] != nil {
				params.Size, _ = strconv.Atoi(query.Get("size"))
			}
			if query["nframes"] != nil {
				params.Nframes, _ = strconv.Atoi(query.Get("nframes"))
			}
			if query["delay"] != nil {
				params.Delay, _ = strconv.Atoi(query.Get("delay"))
			}
			lissajous(w, params)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, defaultLissajousParams)
}

func lissajous(out io.Writer, params LissajousParams) {
	var (
		cycles  = float64(params.Cycles)
		res     = params.Res
		size    = params.Size
		sizef   = float64(params.Size)
		nframes = params.Nframes
		delay   = params.Delay
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*sizef+0.5), size+int(y*sizef+0.5),
				uint8(i % len(palette)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
