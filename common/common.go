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
	//bg, err := colorful.Hex(tm.Bg1)
	//if err != nil {
	//	return err
	//}
	//fg, err := colorful.Hex(tm.Fg1)
	//if err != nil {
	//	return err
	//}
	//var bg01, bg2, bg3, bg4, fg2 string

	if tm.Bg1.HasDarkBG() {
		bg01 = darken(&bg, 0.1)
		bg2 = lighten(&bg, 0.08)
		bg3 = lighten(&bg, 0.16)
		bg4 = lighten(&bg, 0.24)
		fg2 = darken(&fg, 0.08)
	} else {
		bg01 = lighten(&bg, 0.1)
		bg2 = darken(&bg, 0.08)
		bg3 = darken(&bg, 0.16)
		bg4 = darken(&bg, 0.24)
		fg2 = lighten(&fg, 0.08)
	}
	tm.Bg01 = bg01
	tm.DarkBG = hasDarkBG(&bg)
	tm.Bg2 = bg2
	tm.Bg3 = bg3
	tm.Bg4 = bg4
	tm.Fg2 = fg2
	builtin, _ := colorful.Hex(tm.Builtin)
	keyw, _ := colorful.Hex(tm.Keyword)
	typ, _ := colorful.Hex(tm.Type)
	fnc, _ := colorful.Hex(tm.Func)
	warn1, _ := colorful.Hex(tm.Warning)
	warn2, _ := colorful.Hex(tm.Warning2)
	str, _ := colorful.Hex(tm.String)
	tm.InvBuiltin = invertColor(&bg, &builtin)
	tm.InvKeyword = invertColor(&bg, &keyw)
	tm.InvType = invertColor(&bg, &typ)
	tm.InvFunc = invertColor(&bg, &fnc)
	tm.InvString = invertColor(&bg, &str)
	tm.InvWarning = invertColor(&bg, &warn1)
	tm.InvWarning2 = invertColor(&bg, &warn2)
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
