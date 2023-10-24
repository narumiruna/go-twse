package twse

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"encoding/json"

	"github.com/dustin/go-humanize"
	log "github.com/sirupsen/logrus"
)

type Number string

func (n *Number) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*n = Number(s)
	return nil
}

func (n Number) Float64() float64 {
	if n == "-" {
		return 0
	}

	f, err := strconv.ParseFloat(string(n), 64)
	if err != nil {
		return 0
	}

	return f
}

func (n Number) Int64() int64 {
	if n == "-" {
		return 0
	}

	i, err := strconv.ParseInt(string(n), 10, 64)
	if err != nil {
		return 0
	}

	return i
}

type Numbers []Number

func (o *Numbers) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	for _, v := range strings.Split(s, "_") {
		if v == "" {
			continue
		}
		*o = append(*o, Number(v))
	}
	return nil
}

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*t = Time(time.Unix(v/1000, 0))
	return nil
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

type Response struct {
	MsgArray    []StockInfo `json:"msgArray"`
	Referer     string      `json:"referer"`
	UserDelay   int         `json:"userDelay"`
	Rtcode      string      `json:"rtcode"`
	QueryTime   QueryTime   `json:"queryTime"`
	Rtmessage   string      `json:"rtmessage"`
	ExKey       string      `json:"exKey"`
	CachedAlive int         `json:"cachedAlive"`
}

type QueryTime struct {
	SysDate           string `json:"sysDate"`
	StockInfoItem     int    `json:"stockInfoItem"`
	StockInfo         int    `json:"stockInfo"`
	SessionStr        string `json:"sessionStr"`
	SysTime           string `json:"sysTime"`
	ShowChart         bool   `json:"showChart"`
	SessionFromTime   int    `json:"sessionFromTime"`
	SessionLatestTime int    `json:"sessionLatestTime"`
}

type StockInfo struct {
	Symbol            string  `json:"c"`
	Ticker            string  `json:"ch"`
	Name              string  `json:"nf"`
	ShortName         string  `json:"n"`
	Exchange          string  `json:"ex"` // 上市或上櫃
	Ask               Number  `json:"oa"` // 賣價
	Bid               Number  `json:"ob"` // 買價
	Asks              Numbers `json:"a"`  // 五檔賣出價格
	Bids              Numbers `json:"b"`  // 五檔買入價格
	AskVolumes        Numbers `json:"f"`  // 五檔賣出數量
	BidVolumes        Numbers `json:"g"`  // 五檔買入數量
	PrevClose         Number  `json:"y"`  // 昨收
	Open              Number  `json:"o"`  // 開盤
	High              Number  `json:"h"`  // 最高
	Low               Number  `json:"l"`  // 最低
	Close             Number  `json:"z"`  // 收盤
	TradePrice        Number  `json:"pz"` // 成交價
	Volume            Number  `json:"tv"` // 成交量
	AccumulatedVolume Number  `json:"v"`  // 累積成交量
	UpperBound        Number  `json:"u"`  // 漲停價
	LowerBound        Number  `json:"w"`  // 跌停價
	TradeTime         string  `json:"t"`  // 交易時間
	TradeDate         string  `json:"d"`  // 交易日期
	Timestamp         Time    `json:"tlong"`
	Ps                Number  `json:"ps"`
	Nu                string  `json:"nu"` // 網址
	Bp                Number  `json:"bp"`
	Fv                Number  `json:"fv"`
	Ot                string  `json:"ot"` // 某個時間
	IP                Number  `json:"ip"`
	Mt                string  `json:"mt"`
	Ov                Number  `json:"ov"`
	It                Number  `json:"it"`
	Oz                Number  `json:"oz"`
	P                 Number  `json:"p"`
	S                 Number  `json:"s"`
	Ts                Number  `json:"ts"`
}

func (i StockInfo) MidPrice() float64 {
	ask := i.Asks[0].Float64()
	bid := i.Bids[0].Float64()
	return (ask + bid) / 2.0
}

func (i StockInfo) String() string {
	curPrice := i.TradePrice.Float64()
	if curPrice == 0 {
		log.Infof("tradePrice is 0, use midPrice instead")
		curPrice = i.MidPrice()
	}
	netChange := (curPrice/i.PrevClose.Float64() - 1.0) * 100
	return fmt.Sprintf("%s(%s), Open: %s, High: %s, Low: %s, Last: %s, Net Change: %.2f%%, Volume: %d",
		i.ShortName,
		i.Symbol,
		humanize.Commaf(i.Open.Float64()),
		humanize.Commaf(i.High.Float64()),
		humanize.Commaf(i.Low.Float64()),
		humanize.Commaf(i.Close.Float64()),
		netChange,
		i.AccumulatedVolume.Int64(),
	)
}
