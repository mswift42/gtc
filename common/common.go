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

func (tm *ThemeMap) AddColors() error {
	var bg01, bg2, bg3, bg4, fg2 ThemeColor

	if tm.Bg1.HasDarkBG() {
		bg01 = tm.Bg1.Darken(0.1)
		bg2 = tm.Bg1.Lighten(0.08)
		bg3 = tm.Bg1.Lighten(0.16)
		bg4 = tm.Bg1.Lighten(0.24)
		fg2 = tm.Bg1.Darken(0.08)
	} else {
		bg01 = tm.Bg1.Lighten(0.1)
		bg2 = tm.Bg1.Darken(0.08)
		bg3 = tm.Bg1.Darken(0.16)
		bg4 = tm.Bg1.Darken(0.24)
		fg2 = tm.Bg1.Lighten(0.08)
	}
	tm.Bg01 = bg01
	tm.DarkBG = tm.Bg1.HasDarkBG()
	tm.Bg2 = bg2
	tm.Bg3 = bg3
	tm.Bg4 = bg4
	tm.Fg2 = fg2
	tm.InvBuiltin = tm.Builtin.invertColor(bg01)
	tm.InvKeyword = tm.Keyword.invertColor(bg01)
	tm.InvType = tm.Type.invertColor(bg01)
	tm.InvFunc = tm.Func.invertColor(bg01)
	tm.InvString = tm.String.invertColor(bg01)
	tm.InvWarning = tm.Warning.invertColor(bg01)
	tm.InvWarning2 = tm.Warning2.invertColor(bg01)
	return nil
}

type ThemeColor struct {
	col colorful.Color
}

func (tc *ThemeColor) HasDarkBG() bool {
	l, _, _ := tc.col.Lab()
	return l < 0.5
}

func (tc *ThemeColor) Lighten(factor float64) ThemeColor {
	white, _ := colorful.Hex("#ffffff")
	return ThemeColor{tc.col.BlendLab(white, factor)}
}

func (tc *ThemeColor) Darken(factor float64) ThemeColor {
	black, _ := colorful.Hex("#000000")
	return ThemeColor{tc.col.BlendLab(black, factor)}
}

func (tc *ThemeColor) invertColor(bgcol ThemeColor) ThemeColor {
	_, _, l1 := bgcol.col.Hsl()
	h2, s2, _ := tc.col.Hsl()
	return ThemeColor{colorful.Hsl(h2, s2, l1)}
}
