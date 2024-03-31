package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type ExchangeResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
	Data   struct {
		Price float64 `json:"price,string"`
	} `json:"result"`
	exchange string
}

func (e *ExchangeResponse) GetName() string {
	return fmt.Sprintf("%s_response", e.exchange)
}

type RequestDataStep struct {
	Exchange string
	Endpoint string
}

func (s *RequestDataStep) Run(data map[string]daedalus.Data) []daedalus.Data {
	// Send HTTP GET request to the exchange's API.
	resp, err := http.Get(s.Endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON response.
	var exchange_response ExchangeResponse
	err = json.Unmarshal(body, &exchange_response)
	if err != nil {
		panic(err)
	}

	exchange_response.exchange = s.Exchange

	if exchange_response.Data.Price != 0 {
		exchange_response.Price = exchange_response.Data.Price
	}

	// Print the Bitcoin price.
	fmt.Println(s.Exchange, "price:", exchange_response.Price)

	return []daedalus.Data{&exchange_response}
}

func (s *RequestDataStep) GetRequiredData() []string {
	return []string{}
}

func (s *RequestDataStep) GetOutputData() []string {
	return []string{fmt.Sprintf("%s_response", s.Exchange)}
}

type BestPriceData struct {
	Exchange string
	Price    float64
}

func (b *BestPriceData) GetName() string {
	return "BestPrice"
}

type BestPriceStep struct {
}

func (s *BestPriceStep) Run(data map[string]daedalus.Data) []daedalus.Data {
	binance_response := data["Binance_response"].(*ExchangeResponse)
	bybit_response := data["Bybit_response"].(*ExchangeResponse)

	best_price := BestPriceData{Exchange: binance_response.exchange, Price: binance_response.Price}
	if bybit_response.Price < best_price.Price {
		best_price.Exchange = bybit_response.exchange
		best_price.Price = bybit_response.Price
	}

	return []daedalus.Data{&best_price}
}

func (s *BestPriceStep) GetRequiredData() []string {
	return []string{"Binance_response", "Bybit_response"}
}

func (s *BestPriceStep) GetOutputData() []string {
	return []string{"BestPrice"}
}

type PrintBestPriceStep struct {
}

func (s *PrintBestPriceStep) Run(data map[string]daedalus.Data) []daedalus.Data {
	best_price := data["BestPrice"].(*BestPriceData)

	fmt.Println("Best price:", best_price.Exchange, best_price.Price)

	return nil
}

func (s *PrintBestPriceStep) GetRequiredData() []string {
	return []string{"BestPrice"}
}

func (s *PrintBestPriceStep) GetOutputData() []string {
	return []string{}
}

func main() {
	d := daedalus.NewDaedalus()

	// First stage - get the Bitcoin price from Binance and Bybit.
	stage_id, _ := d.AddStep(-1, &RequestDataStep{Exchange: "Binance", Endpoint: "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT"}, true)
	d.AddStep(stage_id, &RequestDataStep{Exchange: "Bybit", Endpoint: "https://api.bybit.com/spot/v3/public/quote/ticker/price?symbol=BTCUSDT"})

	// Second stage - get the best price and print it.
	stage_id, _ = d.AddStep(-1, &BestPriceStep{})
	d.AddStep(stage_id, &PrintBestPriceStep{})

	d.Build()
	d.Run()
}
