package communicate

import (
	"net/http"
	"encoding/json"
	"time"
	"strconv"
)


type KrakenPriceData struct {
	Error []interface{} `json:"error"`
	Result struct {
		XETHZUSD struct {
			A []string `json:"a"`
			B []string `json:"b"`
			C []string `json:"c"`
			V []string `json:"v"`
			P []string `json:"p"`
			T []int `json:"t"`
			L []string `json:"l"`
			H []string `json:"h"`
			O string `json:"o"`
		} `json:"XETHZUSD"`
	} `json:"result"`
}

type GDAXPriceData struct {
	TradeID int `json:"trade_id"`
	Price string `json:"price"`
	Size string `json:"size"`
	Bid string `json:"bid"`
	Ask string `json:"ask"`
	Volume string `json:"volume"`
	Time time.Time `json:"time"`
}

func decodeKrakenDataFromJSON (b []byte) (KrakenPriceData) {
	var decoded KrakenPriceData

	err := json.Unmarshal(b, &decoded)
	if err != nil {
		panic(err)
	}
	return decoded
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
		r, err := httpClient.Get(url)
		if err != nil {
				return err
		}
		defer r.Body.Close()

		return json.NewDecoder(r.Body).Decode(target)
}

func GetPrices() []float32 {

	var decodedKraken KrakenPriceData
	var decodedGDAX GDAXPriceData
	getJson("https://api.kraken.com/0/public/Ticker?pair=ETHUSD", &decodedKraken)
	getJson("https://api.gdax.com/products/ETH-USD/ticker", &decodedGDAX)

	KrakenPrice := decodedKraken.Result.XETHZUSD.P[0]
	GDAXPrice := decodedGDAX.Price

	ConvertedKrakenPrice, err := strconv.ParseFloat(KrakenPrice, 32)
	ConvertedGDAXPrice, err := strconv.ParseFloat(GDAXPrice, 32)
	if (err != nil) {
		panic(err)
	}

	prices := []float32{float32(ConvertedKrakenPrice), float32(ConvertedGDAXPrice)}
	return  prices
}

