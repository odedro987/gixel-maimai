package maimai

import (
	"embed"
	"gixel-maimai/maimai/states"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

//go:embed assets
var assets embed.FS

func init() {
	game := gixel.NewGame(1280, 720, "Maimai", &assets, &states.MenuState{}, 1)

	mobile.SetGame(game)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
