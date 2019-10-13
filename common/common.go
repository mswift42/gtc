package common

import "github.com/lucasb-eyer/go-colorful"

// ThemeMap represents all colors of an ui theme.
// It holds the theme's background color and darker / lighter
// shades thereof, foregroundcolors in varying shades and colors,
// and a boolean that is set to true if the background color counts
// as "dark".
type ThemeMap struct {
	DarkBG      bool
	Fg1         ThemeColor
	Fg2         ThemeColor
	Bg1         ThemeColor
	Bg01        ThemeColor
	Bg2         ThemeColor
	Bg3         ThemeColor
	Bg4         ThemeColor
	Builtin     ThemeColor
	Keyword     ThemeColor
	Constant    ThemeColor
	Comment     ThemeColor
	Func        ThemeColor
	String      ThemeColor
	Type        ThemeColor
	Warning     ThemeColor
	Warning2    ThemeColor
	InvBuiltin  ThemeColor
	InvKeyword  ThemeColor
	InvType     ThemeColor
	InvFunc     ThemeColor
	InvString   ThemeColor
	InvWarning  ThemeColor
	InvWarning2 ThemeColor
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
