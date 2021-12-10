package normal

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	zero.OnCommand("公告", zero.OnlyPrivate, zero.SuperUserPermission).
		Handle(func(ctx *zero.Ctx) {
			groupList := ctx.GetGroupList()
			for _, groupInfo := range groupList.Array() {
				groupId := groupInfo.Get("group_id").Int()
				ctx.SendGroupMessage(groupId, ctx.Event.Message)
			}
		})
}
