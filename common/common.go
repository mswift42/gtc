package common

import "github.com/lucasb-eyer/go-colorful"

// ThemeMap represents all colors of an ui theme.
// It holds the theme's background color and darker / lighter
// shades thereof, foregroundcolors in varying shades and colors,
// and a boolean that is set to true if the background color counts
// as "dark".
type ThemeMap struct {
	DarkBG      bool
	Fg1         string
	Fg2         string
	Bg1         string
	Bg01        string
	Bg2         string
	Bg3         string
	Bg4         string
	Builtin     string
	Keyword     string
	Constant    string
	Comment     string
	Func        string
	String      string
	Type        string
	Warning     string
	Warning2    string
	InvBuiltin  string
	InvKeyword  string
	InvType     string
	InvFunc     string
	InvString   string
	InvWarning  string
	InvWarning2 string
}

type ThemeColor struct {
	col *colorful.Color
}

func (tc *ThemeColor) HasDarkBG() bool {
	l, _, _ := tc.col.Lab()
	return l < 0.5
}

func (tc *ThemeColor) Lighten(factor float64) string {
	white, _ := colorful.Hex("#ffffff")
	return tc.col.BlendLab(white, factor).Hex()
}

func (tc *ThemeColor) Darken(factor float64) string {
	black, _ := colorful.Hex("#000000")
	return tc.col.BlendLab(black, factor).Hex()
}
