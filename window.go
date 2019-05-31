package main

import (
	_ "image"
	_ "image/png"
	_ "os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	img, err := LoadPicture("crying.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(img, img.Bounds())
	win.Clear(colornames.Silver)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	for !win.Closed() {
		win.Update()
	}
}
