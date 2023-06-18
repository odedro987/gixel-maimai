package ui

import (
	"image/color"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/cache"
	"github.com/GixelEngine/gixel-engine/gixel/shader"
)

type Container struct {
	gixel.BaseGxlSprite
}

func NewContainer(x, y float64) *Container {
	c := &Container{}
	c.SetPosition(x, y)
	return c
}

func (c *Container) Init(game *gixel.GxlGame) {
	c.BaseGxlSprite.Init(game)

	c.ApplyGraphic(c.Game().Graphics().MakeGraphic(800, 800, color.Black, cache.CacheOptions{}))
	c.ApplyShader(shader.NewShader("maimai/assets/shaders/container.kage", map[string]interface{}{}))
}
