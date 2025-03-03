package knowledge

import "image/color"

type Person struct {
	Identifier string          `yaml:"identifier"`
	Name       string          `yaml:"name"`
	IconUrl    string          `yaml:"icon_url"`
	ColorRGB   uint32          `yaml:"color"` // 0xRRGGBB
	Twitter    Account[uint64] `yaml:"twitter"`
	Discord    Account[uint64] `yaml:"discord"`
}

func (p Person) Color() color.Color {
	return color.RGBA{
		R: uint8(p.ColorRGB >> 16),
		G: uint8(p.ColorRGB >> 8),
		B: uint8(p.ColorRGB),
		A: 0xFF,
	}
}
