package actions

import (
	"gixel-maimai/maimai/color"
	"gixel-maimai/maimai/shaders"
	ic "image/color"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/cache"
)

type Tap struct {
	Action
	grad color.Gradient2
}

func NewTap(x, y float64, grad color.Gradient2) *Tap {
	t := &Tap{
		grad: grad,
	}
	t.SetPosition(x, y)
	return t
}

func (t *Tap) Init(game *gixel.GxlGame) {
	t.Action.Init(game)
	t.actionType = ActionTap

	//t.ApplyGraphic(game.Graphics().LoadGraphic("maimai/assets/images/tap.png", cache.CacheOptions{}))
	t.ApplyGraphic(t.Game().Graphics().MakeGraphic(100, 100, ic.Black, cache.CacheOptions{}))
	t.ApplyShader(shaders.NewTapShader(t.grad))
}

func (t *Tap) Update(elapsed float64) error {
	t.Shader().Update(elapsed)

	return nil
}
