package shaders

import "github.com/GixelEngine/gixel-engine/gixel/shader"

type TapShader struct {
	shader.BaseGxlShader
}

func NewTapShader(color1 [3]float32, color2 [3]float32) *TapShader {
	return &TapShader{
		BaseGxlShader: *shader.NewShader("maimai/assets/shaders/tap.kage", map[string]interface{}{
			"Time":   float32(0),
			"Color1": color1,
			"Color2": color2,
		}),
	}
}

func (s *TapShader) Update(elapsed float64) {
	s.Uniforms()["Time"] = s.Uniforms()["Time"].(float32) + float32(elapsed)
}
