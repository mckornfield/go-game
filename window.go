package main

import (
	_ "image"
	_ "image/png"
	"math/rand"
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

	var (
		trees    []*pixel.Sprite
		matrices []pixel.Matrix
	)
	for !win.Closed() {
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			tree := getNewTree(treeFrames, spriteSheet)
			trees, matrices = addPlantedTree(trees, matrices, tree, win)
		}

		win.Clear(colornames.Whitesmoke)
		for i, tree := range trees {
			tree.Draw(win, matrices[i])
		}

		win.Update()
	}

}

func addPlantedTree(trees []*pixel.Sprite, matrices []pixel.Matrix, tree *pixel.Sprite, win *pixelgl.Window) ([]*pixel.Sprite, []pixel.Matrix) {
	trees = append(trees, tree)
	matrix := pixel.IM.Scaled(pixel.ZV, 4).Moved(win.MousePosition())
	matrices = append(matrices, matrix)
	return trees, matrices
}

func getNewTree(treeFrames []pixel.Rect, spriteSheet pixel.Picture) *pixel.Sprite {
	treeIndex := rand.Intn(len(treeFrames))
	tree := pixel.NewSprite(spriteSheet, treeFrames[treeIndex])
	return tree
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
