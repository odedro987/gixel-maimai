package color

var (
	Pink = NewGradient2(
		RGBA{255, 0, 120, 255},
		RGBA{255, 160, 230, 255},
	)

	Blue = NewGradient2(
		RGBA{20, 97, 205, 255},
		RGBA{0, 224, 246, 255},
	)

	Double = NewGradient2(
		RGBA{255, 180, 10, 255},
		RGBA{255, 255, 0, 255},
	)

	Break = NewGradient2(
		RGBA{255, 30, 20, 255},
		RGBA{255, 255, 0, 255},
	)
)
