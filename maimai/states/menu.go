package states

import (
	"gixel-maimai/maimai/actions"
	"image/color"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/cache"
)

type MenuState struct {
	gixel.BaseGxlState
}

func (s *MenuState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	bg := gixel.NewSprite(0, 0)
	bg.ApplyGraphic(game.Graphics().MakeGraphic(1280, 800, color.RGBA{255, 0, 255, 255}, cache.CacheOptions{}))
	s.Add(bg)

	s.Add(actions.NewTap(100, 100))

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
