package actions

import (
	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/systems/physics"
)

type ActionType int

const (
	ActionTap ActionType = iota
	ActionHold
	ActionSlide
	ActionTouch
	ActionTouchHold
)

type Action struct {
	// Base
	gixel.BaseGxlSprite
	// Systems
	physics.Physics
	// Action
	actionType ActionType
	dstIdx     int8
}

func (a *Action) Init(game *gixel.GxlGame) {
	a.BaseGxlSprite.Init(game)
	a.Physics.Init(a)
}

func (a *Action) Update(elapsed float64) error {
	err := a.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	a.Physics.Update(elapsed)

	return nil
}
