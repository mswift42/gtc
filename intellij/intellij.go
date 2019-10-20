package intellij

import (
	"encoding/xml"
	"path/filepath"
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

func NewThemeMap(td *ThemeFile) (*common.ThemeMap, error) {
	var tm common.ThemeMap
	am := attrMap(td.ThemeAttrs)
	tm.Bg1 = common.NewThemeColorFromHex(am["TEXT"].bg)
	tm.Fg1 = common.NewThemeColorFromHex(am["TEXT"].fg)
	tm.Func = common.NewThemeColorFromHex(am["DEFAULT_FUNCTION_DECLARATION"].fg)
	tm.Comment = common.NewThemeColorFromHex(am["DEFAULT_BLOCK_COMMENT"].fg)
	tm.Constant = common.NewThemeColorFromHex(am["DEFAULT_CONSTANT"].fg)
	tm.Keyword = common.NewThemeColorFromHex(am["DEFAULT_KEYWORD"].fg)
	tm.String = common.NewThemeColorFromHex(am["DEFAULT_STRING"].fg)
	tm.Type = common.NewThemeColorFromHex(am["DEFAULT_CLASS_NAME"].fg)
	tm.Builtin = common.NewThemeColorFromHex(am["DEFAULT_INSTANCE_FIELD"].fg)
	tm.Warning = common.NewThemeColorFromHex(am["LOG_ERROR_OUTPUT"].fg)
	tm.Warning2 = common.NewThemeColorFromHex(am["LOG_WARNING_OUTPUT"].fg)
	err := tm.AddColors()
	return &tm, err
}

func GenerateTheme(xmlpath, templpath string) error {
	xmlbytes, err := common.LoadFile(xmlpath)
	if err != nil {
		panic(err)
	}
	var tf ThemeFile
	if err := xml.Unmarshal(xmlbytes, &tf); err != nil {
		panic(err)
	}
	tm, err := NewThemeMap(&tf)
	if err != nil {
		panic(err)
	}
	filename := strings.TrimSuffix(xmlpath, filepath.Ext(xmlpath))
	return common.SaveTemplate(templpath, filename, &tm)
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
