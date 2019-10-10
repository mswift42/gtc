package main

import (
	"github.com/lucasb-eyer/go-colorful"
)

func main() {

}

func hasDarkBG(col *colorful.Color) bool {
	l, _, _ := col.Lab()
	return l < 0.5
}

func lighten(col *colorful.Color, factor float64) string {
	white, _ := colorful.Hex("#ffffff")
	return col.BlendLab(white, factor).Hex()
}

func darken(col *colorful.Color, factor float64) string {
	black, _ := colorful.Hex("#000000")
	return col.BlendLab(black, factor).Hex()
}
