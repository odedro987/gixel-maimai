package shaders

import (
	"gixel-maimai/maimai/color"

	"github.com/GixelEngine/gixel-engine/gixel/shader"
)

type TouchShader struct {
	shader.BaseGxlShader
}

func NewTouchShader(gradient color.Gradient2) *TouchShader {
	return &TouchShader{
		BaseGxlShader: *shader.NewShader("maimai/assets/shaders/touch.kage", map[string]interface{}{
			"Time":   float32(0),
			"Color1": gradient.Start().ToRGBSlice(),
			"Color2": gradient.End().ToRGBSlice(),
		}),
	}
}

func (s *TouchShader) Update(elapsed float64) {
	s.Uniforms()["Time"] = s.Uniforms()["Time"].(float32) + float32(elapsed)
}
