package actions

import "github.com/GixelEngine/gixel-engine/gixel"

type ActionType int

const (
	ActionTap ActionType = iota
	ActionHold
	ActionSlide
	ActionTouch
	ActionTouchHold
)

type Action struct {
	gixel.BaseGxlSprite
	actionType ActionType
}

func (a *Action) Init(game *gixel.GxlGame) {
	a.BaseGxlSprite.Init(game)
}
