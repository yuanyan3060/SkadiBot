package group

import (
	"SkadiBot/plugins/utils"
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
	"image"
	"net/http"
)

const decreaseStr = "快走吧，......逃走吧，从这里，从我身边......逃走吧。"

func gerGrayAvatar(UserID int64) (image.Image, error) {
	url := fmt.Sprintf("https://q1.qlogo.cn/g?b=qq&nk=%d&s=100", UserID)
	rep, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(rep.Body)
	if err != nil {
		return nil, err
	}
	return utils.Image2Gray(img), nil
}
func init() {
	zero.OnNotice().SetBlock(false).SetPriority(40).
		Handle(func(ctx *zero.Ctx) {
			if ctx.Event.NoticeType == "group_decrease" {
				//MemberInfo := ctx.GetGroupMemberInfo(ctx.Event.GroupID, ctx.Event.UserID, false)
				//nickname := MemberInfo.Get("nickname").String()
				//text := fmt.Sprintf(decreaseStr, nickname)
				grayAvatar, err := gerGrayAvatar(ctx.Event.UserID)
				if err != nil {
					return
				}
				ctx.SendChain(message.Image("base64://"+helper.BytesToString(utils.Image2Base64(grayAvatar))),
					message.Text(decreaseStr),
					message.At(ctx.Event.UserID),
				)
			}
		})
}
