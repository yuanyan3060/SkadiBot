package normal

import (
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"runtime"
)

func init() {
	zero.OnCommand("占用").
		Handle(func(ctx *zero.Ctx) {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			ctx.SendChain(message.Text(),
				message.Text(fmt.Sprintf("Sys:%d MB\n", (m.Sys-m.HeapReleased)/1024/1024)),
				message.Text(fmt.Sprintf("HeapRetained:%d MB", (m.HeapIdle-m.HeapReleased)/1024/1024)),
			)
		})
}
