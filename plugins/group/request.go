package group

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	zero.OnRequest().SetBlock(false).SetPriority(40).
		Handle(func(ctx *zero.Ctx) {
			if ctx.Event.RequestType == "group" {
				ctx.SetGroupAddRequest(
					ctx.Event.Flag,
					ctx.Event.SubType,
					true,
					"",
				)
			} else if ctx.Event.RequestType == "friend" {
				ctx.SetFriendAddRequest(
					ctx.Event.Flag,
					true,
					"",
				)
			}
		})
}
