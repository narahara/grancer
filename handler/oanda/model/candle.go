package oanda

import "time"

type Candles struct {
	Instrument  string   `json:"instrument"`
	Granularity string   `json:"granularity"`
	Candles     []Candle `json:"candles"`
}

type Candle struct {
	Complete bool      `json:"complete"`
	Volume   int       `json:"volume"`
	Time     time.Time `json:"time"`
	Mid      struct {
		Open  float64 `json:"o,string"`
		High  float64 `json:"h,string"`
		Low   float64 `json:"l,string"`
		Close float64 `json:"c,string"`
	} `json:"mid"`
}
