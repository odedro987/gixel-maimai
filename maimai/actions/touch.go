package actions

import (
	"gixel-maimai/maimai/color"
	"gixel-maimai/maimai/shaders"
	ic "image/color"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/cache"
)

type Touch struct {
	Action
	grad color.Gradient2
}

func NewTouch(x, y float64, grad color.Gradient2) *Touch {
	t := &Touch{
		Action: Action{
			actionType: ActionTouch,
		},
		grad: grad,
	}
	t.SetPosition(x, y)
	return t
}

func (t *Touch) Init(game *gixel.GxlGame) {
	t.Action.Init(game)

	t.ApplyGraphic(t.Game().Graphics().MakeGraphic(160, 160, ic.Black, cache.CacheOptions{}))
	t.ApplyShader(shaders.NewTouchHoldShader())
}

func (t *Touch) Update(elapsed float64) error {
	err := t.Action.Update(elapsed)
	if err != nil {
		return err
	}

	t.Shader().Update(elapsed)

	return nil
}
