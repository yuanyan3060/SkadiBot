package arknights

import (
	"SkadiBot/plugins/utils"
	"encoding/json"
	"fmt"
	"github.com/fogleman/gg"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
	"gorm.io/gorm"
	"image"
	"image/draw"
	"math"
	"math/rand"
	"strings"
	"time"
)

var CharTable map[string]CharData
var Rarity2CharName [][]string

func getProfessionImage(profession string) (image.Image, error) {
	professionImagePath := fmt.Sprintf("data/static/profession_img/%s.png", profession)
	professionImage, err := gg.LoadImage(professionImagePath)
	if err != nil {
		return nil, err
	}
	return professionImage, nil
}

func getRarityImage(rarity int8) (image.Image, error) {
	rarityImagePath := fmt.Sprintf("data/static/gacha_rarity_img/%d.png", rarity)
	rarityImage, err := gg.LoadImage(rarityImagePath)
	if err != nil {
		return nil, err
	}
	return rarityImage, nil
}

func getRarityBackImage(rarity int8, index int) (image.Image, error) {

	rarityImage, err := getRarityImage(rarity)
	if err != nil {
		return nil, err
	}
	rarityBackRGBA := imageToRGBA(rarityImage)
	rarityBackImageRGBA := rarityBackRGBA.SubImage(image.Rect(27+index*123, 0, 149+index*123, 720))
	rarityBackImageCrop := gg.NewContextForImage(rarityBackImageRGBA)
	rarityBackImage := rarityBackImageCrop.Image()
	return rarityBackImage, nil
}

func imageToRGBA(src image.Image) *image.RGBA {
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

func rollGacha() (charNames []string) {
	var rarity int
	for i := 0; i < 10; i++ {
		rarityRand := rand.Float64()
		if rarityRand < 0.02 {
			rarity = 5
		} else if rarityRand < 0.10 {
			rarity = 4
		} else if rarityRand < 0.58 {
			rarity = 3
		} else {
			rarity = 2
		}
		indexRand := rand.Intn(len(Rarity2CharName[rarity]))
		charNames = append(charNames, Rarity2CharName[rarity][indexRand])
	}
	return
}
func gachaTextBuild(charNames []string) (gachaText string) {
	var gachaResult = make([][]string, 6)
	for _, charName := range charNames {
		charData := CharTable[charName]
		gachaResult[charData.Rarity] = append(gachaResult[charData.Rarity], charData.Name)
	}
	for rarity := 5; rarity > 1; rarity-- {
		if len(gachaResult[rarity]) > 0 {
			if rarity == 5 {
				gachaText += "六星干员:"
			} else if rarity == 4 {
				gachaText += "五星干员:"
			} else if rarity == 3 {
				gachaText += "四星干员:"
			} else if rarity == 2 {
				gachaText += "三星干员:"
			}
		}
		for index := 0; index < len(gachaResult[rarity]); index++ {
			if index != len(gachaResult[rarity])-1 {
				gachaText += gachaResult[rarity][index] + ", "
			} else {
				gachaText += gachaResult[rarity][index] + "\n"
			}
		}
	}
	return strings.TrimRight(gachaText, "\n")

}
func drawGachaImage(charNames []string) ([]byte, error) {
	backgroundImage, err := gg.LoadImage("data/static/gacha_background_img/2.png")
	background := gg.NewContextForImage(backgroundImage)
	if err != nil {
		return nil, err
	}
	for index, charName := range charNames {
		charData := CharTable[charName]
		charImagePath := fmt.Sprintf("data/dynamic/char_img/%s.png", charName)
		charImage, err := gg.LoadImage(charImagePath)
		if err != nil {
			return nil, err
		}
		rarityImage, err := getRarityImage(charData.Rarity)
		if err != nil {
			return nil, err
		}
		rarityBackImage, err := getRarityBackImage(charData.Rarity, index)
		if err != nil {
			return nil, err
		}
		professionImage, err := getProfessionImage(charData.Profession)
		if err != nil {
			return nil, err
		}

		background.DrawImage(rarityBackImage, 0, 0)
		background.DrawImage(rarityImage, 27+index*123, 0)
		background.DrawImage(charImage, 27+index*123, 175)
		background.DrawImage(professionImage, 34+int(math.Round(float64(index)*122.5)), 490)
	}
	return utils.Image2Base64(background.Image()), nil
}

func init() {
	zero.OnCommand("十连").
		Handle(func(ctx *zero.Ctx) {
			rollResult := rollGacha()
			i, err := drawGachaImage(rollResult)
			if err != nil {
				return
			}
			sendBase64 := "base64://" + helper.BytesToString(i)

			var user User
			err = DB.First(&user, "qq = ?", ctx.Event.UserID).Error
			if err == gorm.ErrRecordNotFound {
				user = User{
					QQ:    ctx.Event.UserID,
					Chars: "{}",
				}
				DB.Save(&user)
			}
			oldCharDict := make(map[string]int)
			err = json.Unmarshal([]byte(user.Chars), &oldCharDict)
			if err != nil {
				return
			}
			for _, rollChar := range rollResult {
				count, ok := oldCharDict[rollChar]
				if !ok {
					oldCharDict[rollChar] = 1
				} else {
					oldCharDict[rollChar] = count + 1
				}
			}
			CharsBytes, err := json.Marshal(&oldCharDict)
			if err != nil {
				return
			}
			start := time.Now()
			DB.Model(&user).Where("qq = ?", ctx.Event.UserID).Update("chars", string(CharsBytes))
			fmt.Println(time.Now().Sub(start))
			ctx.SendChain(
				message.Image(sendBase64),
				message.Text(gachaTextBuild(rollResult)),
				message.At(ctx.Event.UserID),
			)
		})
}
