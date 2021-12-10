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

const helpStr = "[1]  签到:每日可签到获取十连券\n" +
	"[2]  十连:消耗寻访凭证进行一次十连抽卡\n" +
	"[3]  查询:查询已有的干员\n" +
	"[4]  查公招:根据截图识别标签并计算结果\n" +
	"[5]  源码:获取斯卡蒂bot的仓库\n" +
	"[6]  占用:获取斯卡蒂bot占用的内存\n" +
	"[7]  公告:向斯卡蒂bot加入的群发送一条消息\n" +
	"[8]  群欢迎\n" +
	"[9]  退群提醒\n" +
	"[10] 防撤回（仅发送给超级管理员权限的帐号)\n" +
	"[11] 识别b站视频小程序内的链接"

func drawHelp() (image.Image, error) {
	w := 0.0
	h := utils.Fonts.Metrics().Height / 64
	totalW := 500.0
	totalH := h
	for _, char := range helpStr {
		charWeight := float64(font.MeasureString(utils.Fonts, string(char)) >> 6)
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
