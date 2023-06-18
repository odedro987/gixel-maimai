package shaders

import (
	"gixel-maimai/maimai/color"

	"github.com/GixelEngine/gixel-engine/gixel/shader"
)

type TapShader struct {
	shader.BaseGxlShader
}

func NewTapShader(gradient color.Gradient2) *TapShader {
	return &TapShader{
		BaseGxlShader: *shader.NewShader("maimai/assets/shaders/tap.kage", map[string]interface{}{
			"Time":   float32(0),
			"Color1": gradient.Start().ToRGBSlice(),
			"Color2": gradient.End().ToRGBSlice(),
		}),
	}
}

func (s *TapShader) Update(elapsed float64) {
	s.Uniforms()["Time"] = s.Uniforms()["Time"].(float32) + float32(elapsed)
}
