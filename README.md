# spry

A spritesheet and animation library for [amore](https://github.com/tanema/amore)

Example

```
package main

import (
  "github.com/tanema/amore"
  "github.com/tanema/spry"
)

var (
  playerAtlas *spry.Atlas
  playerRun   *spry.Animation
)

func main() {
  playerAtlas, _ = spry.NewAtlas("images/player.png", 32, 64, 0, 0, 0)
  playerRun = playerAtlas.NewAnimation(
    []int{8, 9, 10, 11}, // frame indexes
    0.5, // animation duration
    true, // this animation loops
  )

  amore.Start(update, draw)
}

func update(dt float32) {
  playerRun.Update(dt)
}

func draw() {
  playerRun.Draw(100, 200) // draw sprite at x: 100, y: 200
}
```
