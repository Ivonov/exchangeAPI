package exchange

import "golang.org/x/net/websocket"
import "fmt"

//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
//-- https://bitfinex.readme.io/v2/docs/open-source-libraries
//-- https://github.com/bitfinexcom/bitfinex-api-go/blob/master/v1/websocket.go
//-- https://docs.bitfinex.com/v2/reference#ws-auth-position
//
//-- https://godoc.org/golang.org/x/net/websocket#pkg-examples
//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
const (
	BTCUSD = "BTCUSD"
	LTCUSD = "LTCUSD"
	LTCBTC = "LTCBTC"
	ETHUSD = "ETHUSD"
	ETHBTC = "ETHBTC"
	ETCUSD = "ETCUSD"
	ETCBTC = "ETCBTC"
	BFXUSD = "BFXUSD"
	BFXBTC = "BFXBTC"
	ZECUSD = "ZECUSD"
	ZECBTC = "ZECBTC"
	XMRUSD = "XMRUSD"
	XMRBTC = "XMRBTC"
	RRTUSD = "RRTUSD"
	RRTBTC = "RRTBTC"
	XRPUSD = "XRPUSD"
	XRPBTC = "XRPBTC"
	EOSETH = "EOSETH"
	EOSUSD = "EOSUSD"
	EOSBTC = "EOSBTC"
	IOTUSD = "IOTUSD"
	IOTBTC = "IOTBTC"
	IOTETH = "IOTETH"
	BCCBTC = "BCCBTC"
	BCUBTC = "BCUBTC"
	BCCUSD = "BCCUSD"
	BCUUSD = "BCUUSD"

	ChanBook   = "book"
	ChanTrade  = "trades"
	ChanTicker = "ticker"

	Trading = "t"
	Funding = "f"
)

//Service is the struct the contains all information of the subscriptions
type Service struct {
	Ws   *websocket.Conn
	subs []Subscriber
}

//Subscriber TODO Trading and f.... ???
type Subscriber struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Symbol  string `json:"symbol"`
}

//PairResponse is the response given by the subscribed pair websocket
//*t == trading
//*f == funding
type PairResponse struct {
	Frr             float32 `json:"FRR"`               //f 	Flash Return Rate - average of all fixed rate funding over the last hour
	Bid             float32 `json:"BID"`               //ft	Price of last highest bid
	BidPeriod       int     `json:"BID_PERIOD"`        //f 	Bid period covered in days
	BidSize         float32 `json:"BID_SIZE"`          //ft Size of the last highest bid
	Ask             float32 `json:"ASK"`               //ft Price of last lowest ask
	AskPeriod       int     `json:"ASK_PERIOD"`        //f	Ask period covered in days
	AskSize         float32 `json:"ASK_SIZE"`          //ft	Size of the last lowest ask
	DailyChange     float32 `json:"DAILY_CHANGE"`      //ft	Amount that the last price has changed since yesterday
	DailyChangePerc float32 `json:"DAILY_CHANGE_PERC"` //ft Amount that the price has changed expressed in percentage terms
	LastPrice       float32 `json:"LAST_PRICE"`        //ft	Price of the last trade
	Volume          float32 `json:"VOLUME"`            //ft	Daily volume
	High            float32 `json:"HIGH"`              //ft Daily high
	Low             float32 `json:"LOW"`               //ft Daily low
}

type websocketInfo struct {
	Event     string  `json:"event"`
	Code      string  `json:"CODE"`
	Message   string  `json:"MSG"`
	Version   float32 `json:"version"`
	Pair      string  `json:"pair"`
	Channel   string  `json:"channel"`
	ChannelID int     `json:"chanId"`
	Symbol    string  `json:"symbol"`
}

// Key of A: A, D, E major G, C, D minor Am Diminished
// {"event":"subscribed","channel":"ticker","chanId":3,"symbol":"tBTCUSD","pair":"BTCUSD"}

type websocketFirstResponse struct {
	Event   string  `json:"event"`
	Version float32 `json:"version"`
}

//Dial TODO ...
func (service *Service) Dial(url, protocol, origin string) (err error) {
	var received websocketInfo
	var ws *websocket.Conn
	ws, err = websocket.Dial(url, protocol, origin)
	if err != nil {
		return err
	}
	service.Ws = ws
	websocket.JSON.Receive(ws, &received)
	fmt.Printf("---- Received from websocket ----\nEvent %s\nVersion %f\n\n", received.Event, received.Version)
	return err
}

//Subscribe TODO ...
func (service *Service) Subscribe(sub Subscriber) (err error) {
	var received websocketInfo
	err = websocket.JSON.Send(service.Ws, sub)
	if err != nil {
		return err
	}
	websocket.JSON.Receive(service.Ws, &received)
	fmt.Printf("Response: %#v\n", received)
	return err
}

// Error codes TODO When connecting check theses codes
// 10000 : Unknown event
// 10001 : Unknown pair

//Websocket info TODO When getting a response check for this info
// {
//    "event":"info",
//    "code": CODE,
//    "msg": MSG
// }
//Websocket info codes TODO
// 20051 : Stop/Restart Websocket Server (please reconnect)
// 20060 : Entering in Maintenance mode. Please pause any activity and resume after receiving the info message 20061 (it should take 120 seconds at most).
// 20061 : Maintenance ended. You can resume normal activity. It is advised to unsubscribe/subscribe again all channels.

//Subscribe
// // request TODO
// {
//    "event": "subscribe",
//    "channel": CHANNEL_NAME
// }
//
// // response TODO
// {
//    "event": "subscribed",
//    "channel": CHANNEL_NAME,
//    "chanId": CHANNEL_ID
// }
//
// // response-failure TODO
// {
//    "event": "error",
//    "msg": ERROR_MSG,
//    "code": ERROR_CODE
// }

//Heartbeat TODO When reseiving info check if the heart is still going and make sure it doesnt create an error expecting a normal response
// [ CHANNEL_ID, "hb" ]

// unsubscribe stuff
// // request TODO
// {
//    "event": "unsubscribe",
//    "chanId": CHANNEL_ID
// }
//
// // response TODO
// {
//    "event": "unsubscribed",
//    "status": "OK",
//    "chanId": CHANNEL_ID
// }
//
// // response-failure TODO
// {
//    "event": "error",
//    "msg": ERROR_MSG,
//    "code": ERROR_CODE
// }
