package intellij

import (
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/mswift42/gtc/common"
)

type AttrOption struct {
	Option string      `xml:"name,attr"`
	Values []AttrValue `xml:"value>option"`
}
type AttrValue struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type ColorOptions struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type ThemeFile struct {
	Colors     []ColorOptions `xml:"colors>option"`
	ThemeAttrs []AttrOption   `xml:"attributes>option"`
}
type themeAttributes struct {
	fg string
	bg string
}

func NewThemeMap(td *ThemeFile) (common.Thememap, error) {
	var tm common.ThemeMap
	am := attrMap(td.ThemeAttrs)
	tm.Bg1 = am["TEXT"].bg
	tm.Fg1 = am["TEXT"].fg
	tm.Func = am["DEFAULT_FUNCTION_DECLARATION"].fg
	tm.Comment = am["DEFAULT_BLOCK_COMMENT"].fg
	tm.Constant = am["DEFAULT_CONSTANT"].fg
	tm.Keyword = am["DEFAULT_KEYWORD"].fg
	tm.String = am["DEFAULT_STRING"].fg
	tm.Type = am["DEFAULT_CLASS_NAME"].fg
	tm.Builtin = am["DEFAULT_INSTANCE_FIELD"].fg
	tm.Warning = am["LOG_ERROR_OUTPUT"].fg
	tm.Warning2 = am["LOG_WARNING_OUTPUT"].fg
	err := tm.addColors()
	return tm, err
}

func attrMap(attros []AttrOption) map[string]themeAttributes {
	tamap := make(map[string]themeAttributes)
	for _, i := range attros {
		var ta themeAttributes
		for _, j := range i.Values {
			lower := strings.ToLower(j.Name)
			if lower == "foreground" {
				ta.fg = j.Value
			}
			if lower == "background" {
				ta.bg = j.Value
			}
		}
		tamap[i.Option] = ta
	}
	return tamap
}

const HexHash = "#"

func (tm *ThemeMap) addColors() error {
	bg, err := colorful.Hex(tm.Bg1)
	if err != nil {
		return err
	}
	fg, err := colorful.Hex(tm.Fg1)
	if err != nil {
		return err
	}
	var bg01, bg2, bg3, bg4, fg2 string

	if hasDarkBG(&bg) {
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
