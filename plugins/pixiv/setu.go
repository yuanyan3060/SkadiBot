package pixiv

import (
	"encoding/json"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/ioutil"
	"net/http"
	"strings"
)

type Setu struct {
	Error string `json:"error"`
	Data  []struct {
		Pid        int      `json:"pid"`
		P          int      `json:"p"`
		Uid        int      `json:"uid"`
		Title      string   `json:"title"`
		Author     string   `json:"author"`
		R18        bool     `json:"r18"`
		Width      int      `json:"width"`
		Height     int      `json:"height"`
		Tags       []string `json:"tags"`
		Ext        string   `json:"ext"`
		UploadDate int64    `json:"uploadDate"`
		Urls       struct {
			Original string `json:"original"`
		} `json:"urls"`
	} `json:"data"`
}

func rollSetu() (setu Setu, err error) {
	get, err := http.Get("https://api.lolicon.app/setu/v2")
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &setu)
	if err != nil {
		return
	}
	return
}

func getImage(url string) ([]byte, error) {
	imgUrl := strings.ReplaceAll(url, "i.pixiv.cat", "i.pixiv.re")
	resp, err := http.Get(imgUrl)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return append(body, []byte{11, 45, 14}...), err
}
func init() {
	zero.OnCommand("涩涩", zero.SuperUserPermission).Handle(func(ctx *zero.Ctx) {
		setu, err := rollSetu()
		if err != nil {
			return
		}
		ctx.SendChain(
			message.Image(strings.ReplaceAll(setu.Data[0].Urls.Original, "i.pixiv.cat", "i.pixiv.re")),
			message.Text(setu.Data[0].Title),
		)
	})
}
