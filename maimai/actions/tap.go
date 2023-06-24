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

func NewTap(x, y float64, grad color.Gradient2, dstIdx int8) *Tap {
	t := &Tap{
		Action: Action{
			actionType: ActionTap,
			dstIdx:     dstIdx,
		},
		grad: grad,
	}
	t.SetPosition(x, y)
	return t
}

func (t *Tap) Init(game *gixel.GxlGame) {
	t.Action.Init(game)

	t.Velocity().Set(100, 100)
	t.Velocity().SetRadians((3.14 / 8) + (6.28/8)*float64(int(t.dstIdx)))

	//t.ApplyGraphic(game.Graphics().LoadGraphic("maimai/assets/images/tap.png", cache.CacheOptions{}))
	t.ApplyGraphic(t.Game().Graphics().MakeGraphic(80, 80, ic.Black, cache.CacheOptions{}))
	t.ApplyShader(shaders.NewTapShader(t.grad))
}

func (t *Tap) Update(elapsed float64) error {
	err := t.Action.Update(elapsed)
	if err != nil {
		return err
	}

	t.Shader().Update(elapsed)

	return nil
}
