package knowledge

import "image/color"

type Person struct {
	Identifier string          `yaml:"identifier"`
	Name       string          `yaml:"name"`
	ColorARGB  uint32          `yaml:"color_argb"` // 0xAARRGGBB make sense?
	Twitter    Account[uint64] `yaml:"twitter"`
	Discord    Account[uint64] `yaml:"discord"`
}

func (p Person) RealColor() color.Color {
	return color.RGBA{
		A: uint8(p.ColorARGB >> 24),
		R: uint8(p.ColorARGB >> 16),
		G: uint8(p.ColorARGB >> 8),
		B: uint8(p.ColorARGB),
	}
}
