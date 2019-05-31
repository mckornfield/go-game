package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func getSprite() *pixel.Sprite {
	img, err := LoadPicture("crying.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(img, img.Bounds())
	return sprite
}

func animateSpinningSprite(win *pixelgl.Window, sprite *pixel.Sprite) {
	win.Clear(colornames.Firebrick)
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
		mat = mat.Moved(win.Bounds().Center())
		sprite.Draw(win, mat)

		win.Update()
	}
}
