# spry [![](https://godoc.org/github.com/tanema/spry?status.svg)](http://godoc.org/github.com/tanema/spry)

A spritesheet and animation library for [amore](https://github.com/tanema/amore) aimed
at being simple and easy to use.

## Example

```golang
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
  image, _ := gfx.NewImage("./explosion.png")
  atlas := spry.NewAtlas(image, 96, 96, 0, 0, 0)
  explosion = atlas.NewAnimation([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 1, true)
  amore.Start(update, draw)
}

func update(dt float32) {
  explosion.Update(dt)
}

func draw() {
  explosion.Draw(3, 3) // draw sprite at x: 3, y: 3
}
```

You can see this example in action in the example folder. You will need to `go get`
amore (and it's prerequisites) to run it.

This demo will change this spritesheet:

![spritesheet](https://raw.githubusercontent.com/tanema/spry/master/example/explosion.png)

into this:

![gif](https://raw.githubusercontent.com/tanema/spry/master/example/explosion.gif)

## Explanation

### Atlas

An atlas is simply an image with a grid of images in it. Each of those images are
a single frame in an animation. So the animation builds a group of quads that are all
the same size.

This is how you create an atlas:

`atlas := spry.NewAtlas(image, frameWidth, frameHeight, left, top, border)`

- `image` is an amore `*gfx.Image`
- `frameWidth` and `frameHeight` are the dimensions of the animation frames. Each
  of the individual "sub-images" that compose the animation. They are usually the
  same size as your character (so if the character is 32x32 pixels, frameWidth is
  32 and so is frameHeight)
- `left` and `top` are the left and top coordinates of the point in the image
  where you want to put the origin of coordinates of the grid. If all the frames
  in your grid are the same size, and the first one's top-left corner is 0,0, then
  you can provide 0, 0 to these. These are often not needed.
- `border` allows you to define "gaps" between your frames in the image. For example,
  imagine that you have frames of 32x32, but they have a 1-px border around each frame.
  So the first frame is not at 0,0, but at 1,1 (because of the border), the second
  one is at 1,33 (for the extra border) etc. You can take this into account and
  "skip" these borders. Pass 0 to border if you don't require it.

### Animation

Animations are groups of frames that are interchanged every now and then.

`animation := atlas.NewAnimation(frames, duration, looping)`

- `frames` is an array of frame indexes. frame indexes start at 0 in the top left and
  increment from left to right. The second row continues where the last row left off.
- `duration` is a float32 describing how long the animation should run for in seconds
- `looping` is a boolean that will determine if the animation should loop when it
  it completes one full animation

For futher methods on animation please see the godoc

## Credits

This was influenced by https://github.com/kikito/anim8 and was altered to be more
in line with Go.
