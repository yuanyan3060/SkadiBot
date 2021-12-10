package utils

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

var Fonts font.Face

func init() {
	var err error
	Fonts, err = gg.LoadFontFace("sarasa-mono-sc-semibold.ttf", 18)
	if err != nil {
		panic(err)
	}
}
