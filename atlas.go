package spry

import (
	"github.com/tanema/amore/gfx"
)

// Atlas is the structure to serialize a spritesheet. You can retrieve and draw
// single frames from the atlas
type Atlas struct {
	sheet  *gfx.Image
	Frames []*gfx.Quad
}

// NewAtlas will build and return a new Atlas. frameWidth and frameHeight are to define
// the size of each frame on the spritesheet. left and top define the offset from
// the top left corner if there is padding around the outside of the spritesheet.
// border defines the amount of pixels inbetween frames.
func NewAtlas(sheet *gfx.Image, frameWidth, frameHeight, left, top, border int32) *Atlas {
	sheetWidth, sheetHeight := sheet.Width, sheet.Height
	framesHorizontal := int32(float32(sheetWidth) / float32(frameWidth))
	framesVertical := int32(float32(sheetHeight) / float32(frameHeight))

	atlas := &Atlas{
		sheet:  sheet,
		Frames: []*gfx.Quad{},
	}

	for y := int32(0); y < framesVertical; y++ {
		for x := int32(0); x < framesHorizontal; x++ {
			atlas.Frames = append(atlas.Frames, gfx.NewQuad(
				left+x*frameWidth+x*border,
				top+y*frameHeight+y*border,
				frameWidth, frameHeight,
				sheetWidth, sheetHeight,
			))
		}
	}

	return atlas
}

// NewAnimation will create and return a new animation using this atlas as reference.
// frames defines the frame indexes in the atlas. duration defines how long the animation
// should play for. looping will define id the animation should loop when it gets
// to the end.
func (atlas *Atlas) NewAnimation(frames []int, duration float32, looping bool) *Animation {
	return &Animation{
		atlas:         atlas,
		frames:        frames,
		duration:      duration,
		frameDuration: duration / float32(len(frames)),
		looping:       looping,
		playing:       true,
	}
}

// GetFrame will return a single frame for drawing
func (atlas *Atlas) GetFrame(index int) *Animation {
	return &Animation{
		atlas:  atlas,
		frames: []int{index},
	}
}

// Draw will draw the frame at the index supplied. the provided arguments fit the
// amore draw interface.
func (atlas *Atlas) Draw(index int, args ...float32) {
	gfx.Drawq(atlas.sheet, atlas.Frames[index], args...)
}
