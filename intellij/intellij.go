package intellij

import (
	"strings"
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
