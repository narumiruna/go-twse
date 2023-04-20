package twse

import (
	"fmt"

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
	TotalVolume int64   `json:"tv,string"`
	Ps          string  `json:"ps"`
	Nu          string  `json:"nu"`
	Pz          string  `json:"pz"`
	Bp          string  `json:"bp"`
	Fv          string  `json:"fv"`
	Oa          string  `json:"oa"`
	Ob          string  `json:"ob"`
	A           string  `json:"a"`
	B           string  `json:"b"`
	Symbol      string  `json:"c"`
	D           string  `json:"d"`
	Ch          string  `json:"ch"`
	Ot          string  `json:"ot"`
	UpdatedAt   string  `json:"tlong"`
	F           string  `json:"f"`
	IP          string  `json:"ip"`
	G           string  `json:"g"`
	Mt          string  `json:"mt"`
	Ov          string  `json:"ov"`
	High        float64 `json:"h,string"`
	It          string  `json:"it"`
	Oz          string  `json:"oz"`
	Low         float64 `json:"l,string"`
	ShortName   string  `json:"n"`
	Open        float64 `json:"o,string"`
	P           string  `json:"p"`
	Ex          string  `json:"ex"`
	S           string  `json:"s"`
	T           string  `json:"t"`
	U           string  `json:"u"`
	Volume      int64   `json:"v,string"`
	W           string  `json:"w"`
	Name        string  `json:"nf"`
	PrevClose   float64 `json:"y,string"`
	Last        float64 `json:"z,string"`
	Ts          string  `json:"ts"`
}

func (i StockInfo) String() string {
	netChange := (i.Last/i.PrevClose - 1.0) * 100
	return fmt.Sprintf("%s(%s), Open: %s, High: %s, Low: %s, Last: %s, Net Change: %f%%",
		i.ShortName,
		i.Symbol,
		humanize.Commaf(i.Open),
		humanize.Commaf(i.High),
		humanize.Commaf(i.Low),
		humanize.Commaf(i.Last),
		netChange,
	)
}
