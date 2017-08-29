package main

import (
	"github.com/Ivonov/exchangeAPI/V2"
	"github.com/golang/glog"
)

func main() {
	origin := "http://localhost/"
	url := "wss://api.bitfinex.com/ws/2"
	service := exchange.Service{}
	err := service.Dial(url, "", origin)
	// service, err := exchange.Dial(url, "", origin)
	if err != nil {
		glog.Error(err)
	}
	defer service.Ws.Close()

	// 	{
	//   "event": "subscribe",
	//   "channel": "ticker",
	//   "symbol": "tBTCUSD"
	// }
	sub := exchange.Subscriber{Event: "subscribe", Channel: "ticker", Symbol: "tBTCUSD"}
	err = service.Subscribe(sub)
	if err != nil {
		glog.Error(err)
	}
}
