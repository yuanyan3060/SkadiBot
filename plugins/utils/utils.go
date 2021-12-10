package utils

import (
	"bytes"
	"encoding/base64"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/jpeg"
	"time"
)

var CstZone = time.FixedZone("CST", 8*3600)

func Image2Base64(image image.Image) []byte {
	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)
	var opt jpeg.Options
	opt.Quality = 95
	_ = jpeg.Encode(encoder, image, &opt)
	err := encoder.Close()
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

func Image2Gray(img image.Image) image.Image {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			colorRgb := img.At(x, y)
			r, g, b, a := colorRgb.RGBA()
			Gray := 0.11*float64(b>>8) + 0.59*float64(g>>8) + 0.3*float64(r>>8)
			var newG uint8
			if Gray >= 255 {
				newG = 255
			} else {
				newG = uint8(Gray)
			}
			newA := uint8(a >> 8)
			// 将每个点的设置为灰度值
			newRgba.SetRGBA(x, y, color.RGBA{R: newG, G: newG, B: newG, A: newA})
		}
	}
	return newRgba
}

func Bytes2Base64(body []byte) string {
	return base64.StdEncoding.EncodeToString(body)
}

func DrawStrImage(str string, W float64) image.Image {
	w := 0.0
	h := Fonts.Metrics().Height / 64
	totalW := W
	totalH := h
	for _, char := range str {
		charWeight := float64(font.MeasureString(Fonts, string(char)) >> 6)
		if w+charWeight > totalW || string(char) == "\n" {
			if string(char) == "\n" {
				w = 0
			} else {
				w = charWeight
			}
			totalH += h
		} else {
			w += charWeight
		}
	}
	var img *gg.Context
	img = gg.NewContext(int(totalW), int(totalH))
	img.SetHexColor("FFFFFF")
	img.Clear()
	img.SetFontFace(Fonts)
	img.SetHexColor("000000")
	w = 0.0
	totalH = h
	for _, char := range str {
		charWeight := float64(font.MeasureString(Fonts, string(char)) >> 6)
		if w+charWeight > totalW || string(char) == "\n" {
			totalH += h
			if string(char) != "\n" {
				w = charWeight
				img.DrawString(string(char), 0.0, float64(totalH))
			} else {
				w = 0
			}
		} else {
			img.DrawString(string(char), w, float64(totalH))
			w += charWeight
		}
	}
	return img.Image()
}
func init() {

}
