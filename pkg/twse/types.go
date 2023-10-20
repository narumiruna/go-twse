package twse

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/narumiruna/go-twse/pkg/types"
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

	Ask        types.Float      `json:"oa"` // 賣價
	Bid        types.Float      `json:"ob"` // 買價
	Asks       types.FloatSlice `json:"a"`  // 五檔賣出價格
	Bids       types.FloatSlice `json:"b"`  // 五檔買入價格
	AskVolumes types.IntSlice   `json:"f"`  // 五檔賣出數量
	BidVolumes types.IntSlice   `json:"g"`  // 五檔買入數量

	PrevClose types.Float `json:"y"` // 昨收
	Open      types.Float `json:"o"` // 開盤
	High      types.Float `json:"h"` // 最高
	Low       types.Float `json:"l"` // 最低
	Close     types.Float `json:"z"` // 收盤

	TradePrice        types.Float `json:"pz"` // 成交價
	Volume            types.Int   `json:"tv"` // 成交量
	AccumulatedVolume types.Int   `json:"v"`  // 累積成交量

	LimitUp   types.Float `json:"u"` // 漲停價
	LimitDown types.Float `json:"w"` // 跌停價

	UpdatedAt string     `json:"t"`
	Timestamp types.Time `json:"tlong"`

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
		humanize.Commaf(float64(i.Open)),
		humanize.Commaf(float64(i.High)),
		humanize.Commaf(float64(i.Low)),
		humanize.Commaf(float64(i.Close)),
		netChange,
		i.AccumulatedVolume,
	)
}
