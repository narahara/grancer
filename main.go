package main

import (
	"encoding/json"

	oandaLib "glancer/handler/oanda/lib"
	oanda "glancer/handler/oanda/model"
	"glancer/lib"
	"log"
)

func main() {

	b, _, err := oandaLib.Get("https://api-fxtrade.oanda.com/v3/instruments/USD_JPY/candles?count=6&price=M&from=2010-02-01T00%3A00%3A00.000000000Z&granularity=H1", "a768709c9055b58ec5ac9055ca406d35-d6440968f9154452ff16dda859788d48")

	if err != nil {
		log.Println(err)
	}
	cc := oanda.Candles{}
	json.Unmarshal(b, &cc)

	lib.PrintJson(cc)

	log.Println(cc.Candles[0].Mid.Open)

}
