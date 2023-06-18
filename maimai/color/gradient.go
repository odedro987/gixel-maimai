package color

import "image/color"

type RGBA color.RGBA

func (c RGBA) ToRGBSlice() [3]float32 {
	return [3]float32{float32(c.R), float32(c.G), float32(c.B)}
}

type Gradient2 struct {
	start RGBA
	end   RGBA
}

func NewGradient2(start, end RGBA) Gradient2 {
	return Gradient2{start: start, end: end}
}

func (g Gradient2) Start() RGBA {
	return g.start
}

func (g Gradient2) End() RGBA {
	return g.end
}
