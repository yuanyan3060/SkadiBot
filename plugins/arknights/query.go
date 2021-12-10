package arknights

import (
	"SkadiBot/plugins/utils"
	"encoding/json"
	"github.com/fogleman/gg"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
	"golang.org/x/image/font"
	"gorm.io/gorm"
	"image"
)

func queryImageBuild(user *User) (image.Image, error) {
	CharDict := make(map[string]int)
	err := json.Unmarshal([]byte(user.Chars), &CharDict)
	if err != nil {
		return nil, err
	}
	totalW := 600.0
	w := 0.0
	h := utils.Fonts.Metrics().Height / 64
	totalH := h
	CharRarityDict := make(map[int8][]string)
	for charId, _ := range CharDict {
		charName := CharTable[charId].Name
		CharRarityDict[CharTable[charId].Rarity] = append(CharRarityDict[CharTable[charId].Rarity], charName)
	}
	for i := int8(5); i > 1; i-- {
		charNames, ok := CharRarityDict[i]
		if ok {
			for _, charName := range charNames {
				charNameWeight := float64(font.MeasureString(utils.Fonts, charName+" ") >> 6)
				if w+charNameWeight > totalW {
					w = charNameWeight
					totalH += h
				} else {
					w += charNameWeight
				}
			}
			totalH += h
			w = 0
		}
	}
	img := gg.NewContext(int(totalW), int(totalH)+5-int(h))
	img.SetHexColor("FFFFFF")
	img.Clear()
	img.SetFontFace(utils.Fonts)
	img.SetHexColor("000000")
	w = 0.0
	h = utils.Fonts.Metrics().Height / 64
	totalH = h
	for i := int8(5); i > 1; i-- {
		charNames, ok := CharRarityDict[i]
		if ok {
			switch i {
			case 5:
				img.SetHexColor("CC3300")
			case 4:
				img.SetHexColor("996600")
			case 3:
				img.SetHexColor("0000CC")
			case 2:
				img.SetHexColor("000000")
			}
			for _, charName := range charNames {
				charNameWeight := float64(font.MeasureString(utils.Fonts, charName+" ") >> 6)
				if w+charNameWeight > totalW {
					totalH += h
					w = charNameWeight
					img.DrawString(charName+" ", 0.0, float64(totalH))
				} else {
					img.DrawString(charName+" ", w, float64(totalH))
					w += charNameWeight
				}
			}
			totalH += h
			w = 0
		}
	}
	return img.Image(), nil
}
func init() {
	var err error

	zero.OnCommand("查询").
		Handle(func(ctx *zero.Ctx) {
			var user User
			err = DB.First(&user, "qq = ?", ctx.Event.UserID).Error
			if err == gorm.ErrRecordNotFound {
				ctx.SendChain(
					message.Text("未找到相关记录"),
				)
			} else {
				img, err := queryImageBuild(&user)
				if err != nil {
					return
				}
				sendBase64 := "base64://" + helper.BytesToString(utils.Image2Base64(img))
				ctx.SendChain(
					message.Image(sendBase64),
				)
			}

		})
}
