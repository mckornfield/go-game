package main

import (
	_ "image"
	_ "image/png"
	_ "os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	win := loadWindow()

	center := win.Bounds().Center()

	win.Clear(colornames.Firebrick)

	sprite := getSprite()

	angle := 0.0
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		win.Clear(colornames.Firebrick)

		mat := pixel.IM
		mat = mat.ScaledXY(pixel.ZV, pixel.V(5, 5))
		mat = mat.Rotated(pixel.ZV, angle)
		angle += 3 * dt
		mat = mat.Moved(center)
		sprite.Draw(win, mat)

		win.Update()
	}
}

func loadWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)
	return win
}

func getSprite() *pixel.Sprite {
	img, err := LoadPicture("crying.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(img, img.Bounds())
	return sprite
}
