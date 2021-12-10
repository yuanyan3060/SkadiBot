package group

import (
	"SkadiBot/plugins/utils"
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/ioutil"
	"net/http"
)

const welcomeStr = "我在等你，博士。我等你太久，太久了，我甚至已经忘了为什么要在这里等你......不过这些都不重要了。不再那么重要了。"

func init() {
	zero.OnNotice().SetBlock(false).FirstPriority().
		Handle(func(ctx *zero.Ctx) {
			if ctx.Event.NoticeType == "group_increase" && ctx.Event.SelfID != ctx.Event.UserID {
				url := fmt.Sprintf("https://q1.qlogo.cn/g?b=qq&nk=%d&s=100", ctx.Event.UserID)
				rep, err := http.Get(url)
				if err != nil {
					return
				}
				all, err := ioutil.ReadAll(rep.Body)
				if err != nil {
					return
				}
				imgBase64 := utils.Bytes2Base64(all)
				ctx.SendChain(message.Image("base64://"+imgBase64),
					message.Text(welcomeStr),
					message.At(ctx.Event.UserID),
				)
			}
		})
}
