package shaders

import (
	"github.com/GixelEngine/gixel-engine/gixel/shader"
)

type TouchHoldShader struct {
	shader.BaseGxlShader
}

func NewTouchHoldShader() *TouchHoldShader {
	return &TouchHoldShader{
		BaseGxlShader: *shader.NewShader("maimai/assets/shaders/touch_hold.kage", map[string]interface{}{
			"Time":     float32(0),
			"Progress": float32(0),
		}),
	}
}

func (s *TouchHoldShader) Update(elapsed float64) {
	s.Uniforms()["Time"] = s.Uniforms()["Time"].(float32) + float32(elapsed)
	if s.Uniforms()["Time"].(float32) >= 1.5 {
		s.Uniforms()["Progress"] = s.Uniforms()["Progress"].(float32) + float32(elapsed)/2
	}
}
