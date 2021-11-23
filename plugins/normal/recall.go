package normal

import (
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strconv"
)

func init() {
	zero.OnNotice().SetBlock(false).SetPriority(40).
		Handle(func(ctx *zero.Ctx) {
			if ctx.Event.NoticeType == "group_recall" || ctx.Event.NoticeType == "friend_recall" {
				recallMessage := ctx.GetMessage(ctx.Event.MessageID)
				for _, SuperUser := range zero.BotConfig.SuperUsers {
					SuperUserId, err := strconv.ParseInt(SuperUser, 10, 64)
					if err != nil {
						return
					}
					var text string
					if ctx.Event.NoticeType == "group_recall" {
						recallUserName := ctx.GetGroupMemberInfo(ctx.Event.GroupID, ctx.Event.UserID, false).Get("nickname").String()
						text = fmt.Sprintf("\n%s(%d)\n%s(%d)\n%s(%d)",
							recallUserName,
							ctx.Event.UserID,
							recallMessage.Sender.NickName,
							recallMessage.Sender.ID,
							ctx.GetGroupInfo(ctx.Event.GroupID, true).Name,
							ctx.Event.GroupID,
						)
					} else {
						text = fmt.Sprintf("\n%s(%d)",
							recallMessage.Sender.NickName,
							recallMessage.Sender.ID,
						)
					}

					sendMessage := append(recallMessage.Elements, message.Text(text))
					ctx.SendPrivateMessage(SuperUserId, sendMessage)
				}

			}
		})
}
