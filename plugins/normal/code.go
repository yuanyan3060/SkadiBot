package normal

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	zero.OnCommand("源码").
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(),
				message.At(ctx.Event.UserID),
			)
		})
}
