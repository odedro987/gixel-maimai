package states

import (
	"gixel-maimai/maimai/actions"
	"gixel-maimai/maimai/color"
	"gixel-maimai/maimai/ui"
	ic "image/color"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/cache"
)

type MenuState struct {
	gixel.BaseGxlState
}

func (s *MenuState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	bg := gixel.NewSprite(0, 0)
	bg.ApplyGraphic(game.Graphics().MakeGraphic(1280, 800, ic.RGBA{255, 0, 255, 255}, cache.CacheOptions{}))
	s.Add(bg)

	container := ui.NewContainer(240, 0)
	s.Add(container)

	for i := 0; i < 8; i++ {
		s.Add(actions.NewTap(640-32, 400-32, color.Pink, int8(i)))
	}

	// for i := 0; i < 1000; i++ {
	// 	s.Add(actions.NewTap(rand.Float64()*700, rand.Float64()*700))
	// }
}

func (s *MenuState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	return nil
}
