package normal

import (
	"SkadiBot/plugins/utils"
	"github.com/fogleman/gg"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
	"golang.org/x/image/font"
	"image"
)

const helpStr = "[1] 十连:消耗寻访凭证进行一次十连抽卡\n" +
	"[2] 查询:查询已有的干员"

func drawHelp() (image.Image, error) {
	w := 0.0
	h := utils.Fonts.Metrics().Height / 64
	totalW := 500.0
	totalH := h
	for _, char := range helpStr {
		charWeight := float64(font.MeasureString(utils.Fonts, string(char)) >> 6)
		if w+charWeight > totalW || string(char) == "\n" {
			w = charWeight
			totalH += h
		} else {
			w += charWeight
		}
	}
	var img *gg.Context
	if int(totalH)+5 > 500 {
		img = gg.NewContext(int(totalW), int(totalH))
	} else {
		img = gg.NewContext(int(totalW), 500)
	}

	img.SetHexColor("FFFFFF")
	img.Clear()
	img.SetFontFace(utils.Fonts)
	img.SetHexColor("000000")
	w = 0.0
	totalH = h
	for _, char := range helpStr {
		charWeight := float64(font.MeasureString(utils.Fonts, string(char)) >> 6)
		if w+charWeight > totalW || string(char) == "\n" {
			w = charWeight
			totalH += h
			img.DrawString(string(char), 0.0, float64(totalH))
		} else {
			img.DrawString(string(char), w, float64(totalH))
			w += charWeight
		}
	}
	return img.Image(), nil
}
func init() {
	zero.OnCommand("help").
		Handle(func(ctx *zero.Ctx) {
			img, err := drawHelp()
			if err != nil {
				return
			}
			sendBase64 := "base64://" + helper.BytesToString(utils.Image2Base64(img))
			ctx.SendChain(
				message.Image(sendBase64),
			)
		})
}
