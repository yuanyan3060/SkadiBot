package normal

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const codeStr = "https://github.com/yuanyan3060/SkadiBot\n人人都可部署的浊心斯卡蒂bot"

func init() {
	zero.OnCommand("源码").
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(codeStr))
		})
}
