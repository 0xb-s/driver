package types

type RGB8 struct {
	R, G, B uint8
}

func NewRGBB(r, g, b uint8) RGB8 {
	return RGB8{R: r, G: g, B: b}
}

type RGBW8 struct {
	R, G, B, W uint8
}

func NewRGBW8(r, g, b, w uint8) RGBW8 {
	return RGBW8{R: r, G: g, B: b, W: w}
}
