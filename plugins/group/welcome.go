package group

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const welcomeStr = "我在等你，博士。我等你太久，太久了，我甚至已经忘了为什么要在这里等你......不过这些都不重要了。不再那么重要了。"

func init() {
	zero.OnNotice().SetBlock(false).FirstPriority().
		Handle(func(ctx *zero.Ctx) {
			if ctx.Event.NoticeType == "group_increase" {
				ctx.SendChain(message.Text(welcomeStr),
					message.At(ctx.Event.UserID),
				)
			}
		})
}
