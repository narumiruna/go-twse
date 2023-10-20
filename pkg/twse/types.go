package twse

import (
	"fmt"
	"strconv"

	"github.com/dustin/go-humanize"
)

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
	Symbol    string `json:"c"`
	Ticker    string `json:"ch"`
	Name      string `json:"nf"`
	ShortName string `json:"n"`
	Exchange  string `json:"ex"` // 上市或上櫃

	Ask        float64 `json:"oa,string"` // 賣價
	Bid        float64 `json:"ob,string"` // 買價
	Asks       string  `json:"a"`         // 五檔賣出價格
	Bids       string  `json:"b"`         // 五檔買入價格
	AskVolumes string  `json:"f"`         // 五檔賣出數量
	BidVolumes string  `json:"g"`         // 五檔買入數量

	PrevClose float64 `json:"y,string"` // 昨收
	Open      float64 `json:"o,string"` // 開盤
	High      float64 `json:"h,string"` // 最高
	Low       float64 `json:"l,string"` // 最低
	Close     float64 `json:"z,string"` // 收盤

	TradePrice        float64 `json:"pz,string"` // 成交價
	Volume            int64   `json:"tv,string"` // 成交量
	AccumulatedVolume int64   `json:"v,string"`  // 累積成交量

	LimitUp   string `json:"u"` // 漲停價
	LimitDown string `json:"w"` // 跌停價

	UpdatedAt string `json:"t"`
	Timestamp int64  `json:"tlong,string"`

	TradeDate string `json:"d"` // 最近交易日期

	Ps string `json:"ps"`
	Nu string `json:"nu"`
	Bp string `json:"bp"`
	Fv string `json:"fv"`
	Ot string `json:"ot"`
	IP string `json:"ip"`
	Mt string `json:"mt"`
	Ov string `json:"ov"`
	It string `json:"it"`
	Oz string `json:"oz"`
	P  string `json:"p"`
	S  string `json:"s"`
	Ts string `json:"ts"`
}

func (i StockInfo) String() string {
	netChange := (i.TradePrice/i.PrevClose - 1.0) * 100
	return fmt.Sprintf("%s(%s), Open: %s, High: %s, Low: %s, Last: %s, Net Change: %.2f%%, Volume: %d",
		i.ShortName,
		i.Symbol,
		humanize.Commaf(i.Open),
		humanize.Commaf(i.High),
		humanize.Commaf(i.Low),
		humanize.Commaf(i.Close),
		netChange,
		i.AccumulatedVolume,
	)
}
