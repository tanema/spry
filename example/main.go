package main

import (
	"github.com/tanema/amore"
	"github.com/tanema/amore/gfx"
	"github.com/tanema/spry"
)

var (
	explosion *spry.Animation
)

func main() {
	image, err := gfx.NewImage("./explosion.png")
	if err != nil {
		panic(err.Error())
	}

	atlas := spry.NewAtlas(image, 96, 96, 0, 0, 0)
	explosion = atlas.NewAnimation(
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, // frame indexes
		1,    // animation duration 1 second
		true, // this animation loops
	)

	amore.Start(update, draw)
}

func update(dt float32) {
	explosion.Update(dt)
}

func draw() {
	explosion.Draw(3, 3) // draw sprite at x: 3, y: 3
}
