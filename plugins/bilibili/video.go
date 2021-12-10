package bilibili

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type BilibiliVideoAppInfo struct {
	App    string `json:"app"`
	Desc   string `json:"desc"`
	View   string `json:"view"`
	Ver    string `json:"ver"`
	Prompt string `json:"prompt"`
	Meta   struct {
		Detail1 struct {
			Appid         string `json:"appid"`
			Desc          string `json:"desc"`
			GamePoints    string `json:"gamePoints"`
			GamePointsUrl string `json:"gamePointsUrl"`
			Host          struct {
				Nick string `json:"nick"`
				Uin  int    `json:"uin"`
			} `json:"host"`
			Icon              string `json:"icon"`
			Preview           string `json:"preview"`
			Qqdocurl          string `json:"qqdocurl"`
			Scene             int    `json:"scene"`
			ShareTemplateData struct {
			} `json:"shareTemplateData"`
			ShareTemplateId string `json:"shareTemplateId"`
			ShowLittleTail  string `json:"showLittleTail"`
			Title           string `json:"title"`
			Url             string `json:"url"`
		} `json:"detail_1"`
	} `json:"meta"`
	Config struct {
		AutoSize int    `json:"autoSize"`
		Ctime    int    `json:"ctime"`
		Forward  int    `json:"forward"`
		Height   int    `json:"height"`
		Token    string `json:"token"`
		Type     string `json:"type"`
		Width    int    `json:"width"`
	} `json:"config"`
}

type VideoData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Bvid      string      `json:"bvid"`
		Aid       int         `json:"aid"`
		Videos    int         `json:"videos"`
		Tid       int         `json:"tid"`
		Tname     string      `json:"tname"`
		Copyright int         `json:"copyright"`
		Pic       string      `json:"pic"`
		Title     string      `json:"title"`
		Pubdate   int         `json:"pubdate"`
		Ctime     int         `json:"ctime"`
		Desc      string      `json:"desc"`
		DescV2    interface{} `json:"desc_v2"`
		State     int         `json:"state"`
		Duration  int         `json:"duration"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			ArgueMsg   string `json:"argue_msg"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		NoCache bool `json:"no_cache"`
		Pages   []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			FirstFrame string `json:"first_frame"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
		UserGarb struct {
			UrlImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
		} `json:"honor_reply"`
	} `json:"data"`
}

func (b *BilibiliVideoAppInfo) getRealUrl() (string, error) {
	for i := 0; i <= 5; i++ {
		rep, err := http.Get(b.Meta.Detail1.Qqdocurl)
		//err = errors.Unwrap(err)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return "", err
			}
		}
		if rep == nil {
			logrus.Debug("get redirect url fail, retry")
		} else {
			url := rep.Request.URL
			return url.Scheme + "://" + url.Host + url.Path, nil
		}
	}
	return "", errors.New("get redirect url fail too many times")
}

func getVideoData(url string) (videoData VideoData, err error) {
	bvRegex, err := regexp.Compile("BV([a-zA-Z0-9]{10})+")
	if err != nil {
		return
	}
	bvId := bvRegex.FindString(url)
	avRegex, err := regexp.Compile("av\\d+")
	if err != nil {
		return
	}
	avId := avRegex.FindString(url)
	var apiUrl string
	if bvId != "" {
		apiUrl = fmt.Sprintf("https://api.bilibili.com/x/web-interface/view?bvid=%s", bvId)
	} else if avId != "" {
		apiUrl = fmt.Sprintf("https://api.bilibili.com/x/web-interface/view?bvid=%s", avId)
	} else {
		err = errors.New("regex fail")
		return
	}
	rep, err := http.Get(apiUrl)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &videoData)
	if err != nil {
		return
	}
	return
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func init() {
	zero.OnMessage().SetBlock(false).FirstPriority().
		Handle(func(ctx *zero.Ctx) {
			for _, msg := range ctx.Event.Message {
				if msg.Type == "json" {
					data, ok := msg.Data["data"]
					if ok {
						var bilibiliVideoAppInfo BilibiliVideoAppInfo
						err := json.Unmarshal([]byte(data), &bilibiliVideoAppInfo)
						if err != nil {
							return
						}
						url, err := bilibiliVideoAppInfo.getRealUrl()
						if err != nil {
							return
						}
						videoData, err := getVideoData(url)
						if err != nil {
							return
						}
						timeStr := time.Unix(int64(videoData.Data.Pubdate), 0).Format("2006 01 02 15:04:05")
						durationStr := fmtDuration(time.Duration(videoData.Data.Duration) * time.Minute)
						ctx.SendChain(message.Text(fmt.Sprintf("视频:%s\n", bilibiliVideoAppInfo.Meta.Detail1.Desc)),
							message.Text(fmt.Sprintf("作者:%s\n", videoData.Data.Owner.Name)),
							message.Text(fmt.Sprintf("时间:%s\n", timeStr)),
							message.Text(fmt.Sprintf("时长:%s\n", durationStr)),
							message.Text(fmt.Sprintf("链接:%s\n", url)),
							message.Image(videoData.Data.Pic),
						)

					}
				}
			}
		})
}
