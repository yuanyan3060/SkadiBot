package arknights

import (
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"gorm.io/gorm"
	"time"
)

func init() {
	zero.OnCommand("签到").
		Handle(func(ctx *zero.Ctx) {
			var user User
			err := DB.First(&user, "qq = ?", ctx.Event.UserID).Error
			if err == gorm.ErrRecordNotFound {
				user = User{
					QQ:              ctx.Event.UserID,
					Chars:           "{}",
					TenGachaTickets: 0,
					LastCheckInTime: time.Unix(0, 0),
				}
			}
			if user.isCanCheckIn() {
				user.LastCheckInTime = time.Now()
				user.TenGachaTickets = 10
				ctx.SendChain(
					message.Text("签到获得10十连券"),
				)
				err = DB.Save(&user).Error
				err := DB.First(&user, "qq = ?", ctx.Event.UserID).Error
				fmt.Println(user.TenGachaTickets)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				ctx.SendChain(
					message.Text("今日已签过到"),
				)
			}
		})
}
