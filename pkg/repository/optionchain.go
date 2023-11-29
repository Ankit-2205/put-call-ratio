package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	NSEBasePath         = "https://www.nseindia.com"
	TradeEntityEndpoint = "/api/option-chain-indices"
)

type TradeEntity struct {
	StrikePrice           float64 `json:"strikePrice"`
	ExpiryDate            string  `json:"expiryDate"`
	Underlying            string  `json:"underlying"`
	Identifier            string  `json:"identifier"`
	OpenInterest          float64 `json:"openInterest"`
	ChangeinOpenInterest  float64 `json:"changeinOpenInterest"`
	PchangeinOpenInterest float64 `json:"pchangeinOpenInterest"`
	TotalTradedVolume     int     `json:"totalTradedVolume"`
	ImpliedVolatility     float64 `json:"impliedVolatility"`
	LastPrice             float64 `json:"lastPrice"`
	Change                float64 `json:"change"`
	PChange               float64 `json:"pChange"`
	TotalBuyQuantity      int     `json:"totalBuyQuantity"`
	TotalSellQuantity     int     `json:"totalSellQuantity"`
	BidQty                int     `json:"bidQty"`
	Bidprice              float64 `json:"bidprice"`
	AskQty                int     `json:"askQty"`
	AskPrice              float64 `json:"askPrice"`
	UnderlyingValue       float64 `json:"underlyingValue"`
}

type RecordData struct {
	StrikePrice float64     `json:"strikePrice"`
	ExpiryDate  string      `json:"expiryDate"`
	PE          TradeEntity `json:"PE"`
	CE          TradeEntity `json:"CE"`
}

type Records struct {
	ExpiryDates     []string     `json:"expiryDates"`
	Data            []RecordData `json:"data"`
	Timestamp       string       `json:"timestamp"`
	UnderlyingValue float64      `json:"underlyingValue"`
}

type OptionChain struct {
	Records Records `json:"records"`
}

type optionChainRepo struct {
	client http.Client
}

func (o *optionChainRepo) GetOptionChain(index string) (OptionChain, error) {

	var optionChain OptionChain

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?symbol=%s", NSEBasePath, TradeEntityEndpoint, index), nil)
	if err != nil {
		return optionChain, err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")

	resp, err := o.client.Do(req)
	if err != nil {
		return optionChain, err
	}

	if resp.StatusCode != 200 {
		return optionChain, errors.New(fmt.Sprintf("Error response code: %d", resp.StatusCode))
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return optionChain, err
	}
	err = json.Unmarshal(respData, &optionChain)

	return optionChain, err
}

func NewOptionChainRepo(client http.Client) Repository {
	return &optionChainRepo{
		client: client,
	}
}
