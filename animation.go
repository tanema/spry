package spry

// Animation defines a range of frames from the atlas to be played over a specified
// amount of time.
type Animation struct {
	atlas         *Atlas
	frames        []int
	duration      float32
	frameDuration float32
	looping       bool
	position      int
	playing       bool
	flippedH      bool
	flippedV      bool
	timer         float32
}

// FlipH will flip the animation horizontally while animating
func (anim *Animation) FlipH() {
	anim.flippedH = !anim.flippedH
}

// FlipV will flip the animation vertically while animating
func (anim *Animation) FlipV() {
	anim.flippedV = !anim.flippedV
}

// Pause will pause the animation on the current frame
func (anim *Animation) Pause() {
	anim.playing = false
}

// IsPlaying will return if the animation is currently animating
func (anim *Animation) IsPlaying() bool {
	return anim.playing
}

// Resume will make the animation continue if it was paused
func (anim *Animation) Resume() {
	anim.playing = true
}

// GotoFrame will forward the animation to the specified position in the animation
func (anim *Animation) GotoFrame(position int) {
	anim.position = position
	anim.timer = float32(position) * anim.frameDuration
}

// PauseAtEnd will forward the animation to the end and pause it
func (anim *Animation) PauseAtEnd() {
	anim.position = len(anim.frames) - 1
	anim.timer = anim.duration
	anim.Pause()
}

// PauseAtStart will return to the beginnign of the animation and pause it.
func (anim *Animation) PauseAtStart() {
	anim.position = 1
	anim.timer = 0
	anim.Pause()
}

// Update will take in the difference in time (dt) and update the current frame
// it is displaying.
func (anim *Animation) Update(dt float32) {
	if !anim.playing {
		return
	}

	anim.timer = anim.timer + dt
	loops := int(anim.timer / anim.duration)
	if loops != 0 {
		if anim.looping {
			anim.timer = anim.timer - anim.duration*float32(loops)
		} else {
			anim.playing = false
			anim.position = len(anim.frames) - 1
			return
		}
	}

	anim.position = int(float32(len(anim.frames)) * (anim.timer / anim.duration))
}

// Draw will draw the current frame of the animation to the screen. The arguments
// satisfy the draw interface from amore
func (anim *Animation) Draw(args ...float32) {
	x, y, angle, sx, sy, ox, oy, kx, ky := anim.normalizeDrawCallArgs(args)
	anim.atlas.Draw(anim.frames[anim.position], x, y, angle, sx, sy, ox, oy, kx, ky)
}

func (anim *Animation) normalizeDrawCallArgs(args []float32) (float32, float32, float32, float32, float32, float32, float32, float32, float32) {
	var x, y, angle, sx, sy, ox, oy, kx, ky float32
	sx = 1
	sy = 1

	argsLength := len(args)

	switch argsLength {
	case 9:
		ky = args[8]
		fallthrough
	case 8:
		kx = args[7]
		if argsLength == 8 {
			ky = kx
		}
		fallthrough
	case 7:
		oy = args[6]
		fallthrough
	case 6:
		ox = args[5]
		if argsLength == 6 {
			oy = ox
		}
		fallthrough
	case 5:
		sy = args[4]
		fallthrough
	case 4:
		sx = args[3]
		if argsLength == 4 {
			sy = sx
		}
		fallthrough
	case 3:
		angle = args[2]
		fallthrough
	case 2:
		x = args[0]
		y = args[1]
	}

	w, h := anim.GetDimensions()
	if anim.flippedH {
		sx = sx * -1
		ox = w - ox
		kx = kx * -1
		ky = ky * -1
	}

	if anim.flippedV {
		sy = sy * -1
		oy = h - oy
		kx = kx * -1
		ky = ky * -1
	}

	return x, y, angle, sx, sy, ox, oy, kx, ky
}

// GetDimensions will return the width and height of the current frame
func (anim *Animation) GetDimensions() (float32, float32) {
	frameIndex := anim.frames[anim.position]
	frame := anim.atlas.Frames[frameIndex]
	_, _, w, h := frame.GetViewport()
	return float32(w), float32(h)
}
