package actions

import (
	"gixel-maimai/maimai/shaders"
	"image/color"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/cache"
)

type Tap struct {
	Action
}

func NewTap(x, y float64) *Tap {
	t := &Tap{}
	t.SetPosition(x, y)
	return t
}

func (t *Tap) Init(game *gixel.GxlGame) {
	t.Action.Init(game)
	t.actionType = ActionTap

	//t.ApplyGraphic(game.Graphics().LoadGraphic("maimai/assets/images/tap.png", cache.CacheOptions{}))
	t.ApplyGraphic(t.Game().Graphics().MakeGraphic(100, 100, color.Black, cache.CacheOptions{}))
	t.ApplyShader(shaders.NewTapShader([3]float32{255, 0, 120}, [3]float32{255, 160, 230}))
}

func (t *Tap) Update(elapsed float64) error {
	t.Shader().Update(elapsed)

	return nil
}
