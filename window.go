package main

import (
	_ "image"
	_ "image/png"
	_ "os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	win := loadWindow()

	sprite := getSprite()

	animateSpinningSprite(win, sprite)

	//spriteSheet := getSpriteSheet()

	//tree := pixel.NewSprite(spriteSheet, pixel.R(0, 0, 32, 32))

	/*for !win.Closed() {
		win.Clear(colornames.Whitesmoke)
		tree.Draw(win, pixel.IM.Scaled(pixel.ZV, 16).Moved(win.Bounds().Center()))
		win.Update()
	}
	*/
}

func loadWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Woot!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}

func getSpriteSheet() pixel.Picture {
	img, err := LoadPicture("trees.png")
	if err != nil {
		panic(err)
	}
	return img
}
