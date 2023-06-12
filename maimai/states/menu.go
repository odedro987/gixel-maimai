package states

import (
	"github.com/GixelEngine/gixel-engine/gixel"
)

type MenuState struct {
	gixel.BaseGxlState
}

func (s *MenuState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)
}

func (s *MenuState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	return nil
}
