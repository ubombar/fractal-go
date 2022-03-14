package main

import (
	"image/color"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func drawTriange(win *pixelgl.Window, a, b, c pixel.Vec, color color.Color) {
	imd := imdraw.New(nil)

	imd.Color = color
	imd.Push(a)
	imd.Push(b)
	imd.Push(c)
	imd.Polygon(0)

	imd.Draw(win)
}

func drawFractal(win *pixelgl.Window, a, b, c pixel.Vec, level int) {
	drawTriange(win, a, b, c, colornames.Orange)

	ab := pixel.Lerp(a, b, 0.5)
	bc := pixel.Lerp(b, c, 0.5)
	ca := pixel.Lerp(c, a, 0.5)

	drawTriange(win, ab, bc, ca, colornames.Skyblue)

	if level > 0 {
		drawFractal(win, ab, b, bc, level-1)
		drawFractal(win, a, ab, ca, level-1)
		drawFractal(win, ca, bc, c, level-1)
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Fractal Test",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	past := time.Now().UnixNano()
	var t float64
	T := 5000.0 * 2

	for !win.Closed() {
		// The delay is in MS
		delay := float64(time.Now().UnixNano()-past) / 1000000.0
		t += delay
		past = time.Now().UnixNano()

		levels_to_draw := (1 + math.Sin(math.Pi*2*t/T)) * 3.6

		win.Clear(colornames.Skyblue)
		drawFractal(win,
			pixel.V(200, 100),
			pixel.V(800, 100),
			pixel.V(500, 700),
			int(levels_to_draw))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
