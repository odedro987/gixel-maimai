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

	s.Add(actions.NewTap(640-50, 100, color.Pink))
	s.Add(actions.NewTap(640+50, 100, color.Blue))
	s.Add(actions.NewTap(640-50, 200, color.Double))
	s.Add(actions.NewTap(640+50, 200, color.Break))

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
