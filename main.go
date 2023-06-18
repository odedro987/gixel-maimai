package main

import (
	"embed"
	"gixel-maimai/maimai/states"

	"github.com/GixelEngine/gixel-engine/gixel"
)

//go:embed maimai/assets
var assets embed.FS

func main() {
	gixel.NewGame(1280, 800, "Maimai", &assets, &states.MenuState{}, 1).Run()
}
