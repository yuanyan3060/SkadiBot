package arknights

import (
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type CharData struct {
	Name               string `json:"name"`
	Profession         string `json:"profession"`
	Rarity             int8   `json:"rarity"`
	ItemObtainApproach string `json:"itemObtainApproach"`
}

type User struct {
	gorm.Model
	QQ              int64 `gorm:"primary_key"`
	Chars           string
	tenGachaTickets int
	lastCheckInTime float64
}

var DB *gorm.DB

func init() {
	var err error

	file, err := ioutil.ReadFile("data/dynamic/gamedata/excel/character_table.json")
	if err != nil {
		return
	}
	Rarity2CharName = make([][]string, 6)
	err = json.Unmarshal(file, &CharTable)
	for charId, chardata := range CharTable {
		if len(chardata.ItemObtainApproach) > 0 {
			Rarity2CharName[chardata.Rarity] = append(Rarity2CharName[chardata.Rarity], charId)
		}
	}

	slowLogger := logger.New(
		//将标准输出作为Writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),

		logger.Config{
			//设定慢查询时间阈值为800ms
			SlowThreshold: 800 * time.Millisecond,
			//设置日志级别，只有Warn和Info级别会输出慢查询日志
			LogLevel: logger.Warn,
		},
	)
	DB, err = gorm.Open(sqlite.Open("arknights.db"), &gorm.Config{
		Logger: slowLogger,
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}
