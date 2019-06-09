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
	win := loadWindow()

	spriteSheet := getSpriteSheet()

	treeFrames := getTreeFrames(spriteSheet)

	tree := pixel.NewSprite(spriteSheet, treeFrames[0])

	for !win.Closed() {
		win.Clear(colornames.Whitesmoke)
		tree.Draw(win, pixel.IM.Scaled(pixel.ZV, 16).Moved(win.Bounds().Center()))
		win.Update()
	}

}

func getTreeFrames(spriteSheet pixel.Picture) []pixel.Rect {
	treeFrames := []pixel.Rect{}
	for x := spriteSheet.Bounds().Min.X; x < spriteSheet.Bounds().Max.X; x += 32 {
		for y := spriteSheet.Bounds().Min.Y; y < spriteSheet.Bounds().Max.Y; y += 32 {
			treeFrames = append(treeFrames, pixel.R(x, y, x+32, y+32))
		}
	}
	return treeFrames
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
