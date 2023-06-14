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
	Volume      string `json:"tv"`
	Ps          string `json:"ps"`
	Nu          string `json:"nu"`
	Pz          string `json:"pz"`
	Bp          string `json:"bp"`
	Fv          string `json:"fv"`
	Oa          string `json:"oa"`
	Ob          string `json:"ob"`
	A           string `json:"a"`
	B           string `json:"b"`
	Symbol      string `json:"c"`
	D           string `json:"d"`
	Ch          string `json:"ch"`
	Ot          string `json:"ot"`
	UpdatedAt   string `json:"tlong"`
	F           string `json:"f"`
	IP          string `json:"ip"`
	G           string `json:"g"`
	Mt          string `json:"mt"`
	Ov          string `json:"ov"`
	High        string `json:"h"`
	It          string `json:"it"`
	Oz          string `json:"oz"`
	Low         string `json:"l"`
	ShortName   string `json:"n"`
	Open        string `json:"o"`
	P           string `json:"p"`
	Ex          string `json:"ex"`
	S           string `json:"s"`
	T           string `json:"t"`
	U           string `json:"u"`
	TotalVolume string `json:"v"`
	W           string `json:"w"`
	Name        string `json:"nf"`
	PrevClose   string `json:"y"`
	Last        string `json:"z"`
	Ts          string `json:"ts"`
}

func (i StockInfo) String() string {
	last, err := strconv.ParseFloat(i.Last, 64)
	if err != nil {
		last = 0
	}

	prevClose, err := strconv.ParseFloat(i.PrevClose, 64)
	if err != nil {
		prevClose = 0
	}

	open, err := strconv.ParseFloat(i.Open, 64)
	if err != nil {
		open = 0
	}

	high, err := strconv.ParseFloat(i.High, 64)
	if err != nil {
		high = 0
	}

	low, err := strconv.ParseFloat(i.Low, 64)
	if err != nil {
		low = 0
	}

	netChange := (last/prevClose - 1.0) * 100
	if last == 0 || prevClose == 0 {
		netChange = 0.0
	}

	return fmt.Sprintf("%s(%s), Open: %s, High: %s, Low: %s, Last: %s, Net Change: %.2f%%, Volume: %s",
		i.ShortName,
		i.Symbol,
		humanize.Commaf(open),
		humanize.Commaf(high),
		humanize.Commaf(low),
		humanize.Commaf(last),
		netChange,
		i.TotalVolume,
	)
}
