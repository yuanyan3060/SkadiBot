package group

import (
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const decreaseStr = "快走吧，%s......逃走吧，从这里，从我身边......逃走吧。"

func init() {
	zero.OnNotice().SetBlock(false).SetPriority(40).
		Handle(func(ctx *zero.Ctx) {
			if ctx.Event.NoticeType == "group_decrease" {
				MemberInfo := ctx.GetGroupMemberInfo(ctx.Event.GroupID, ctx.Event.UserID, false)
				nickname := MemberInfo.Get("nickname").String()
				text := fmt.Sprintf(decreaseStr, nickname)
				ctx.SendChain(message.Text(text))
			}
		})
}
