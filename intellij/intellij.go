package intellij

import "strings"

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

type themeAttributes struct {
	fg string
	bg string
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
