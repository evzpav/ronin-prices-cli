package prices

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var defaultTimeout = 10

//NewClient for prices api
func NewClient(token string) *PricesInput {
	return &PricesInput{
		Source:   "coinmarketcap",
		APIToken: token,
		HTTPClient: &http.Client{
			Timeout: setSecondsDuration(defaultTimeout),
		},
	}
}

//SetPriceSource set data price source
func (p *PricesInput) SetPriceSource(source string) {
	p.Source = source
}

//ChangeHTTPClient set new http client if needed
func (p *PricesInput) ChangeHTTPClient(newHTTPClient *http.Client) {
	p.HTTPClient = newHTTPClient
}

//ChangeTimeout set new timeout in seconds
func (p *PricesInput) ChangeTimeout(newTimeout int) {
	p.HTTPClient.Timeout = setSecondsDuration(newTimeout)
}

//GetCurrencyConvertData reusable function
func (p *PricesInput) GetCurrencyConvertData(uri, symbol, convert string) RoninPricesResp {
	url := fmt.Sprintf("%s/%s?symbol=%s&convert=%s&source=%s", "http://prices.stratum.bt", uri, symbol, convert, p.Source)
	body := p.GetRequest(url, p.APIToken)

	var resp RoninPricesResp
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Printf("Error unmarshal: %s", err.Error())
	}
	return resp
}

//GetCurrency get currency full return
func (p *PricesInput) GetCurrency(symbol, convert string) RoninPricesResp {
	return p.GetCurrencyConvertData("currency", symbol, convert)
}

//GetCurrencies get currencies full return
func (p *PricesInput) GetCurrencies(symbols, converts string) RoninPricesResp {
	return p.GetCurrencyConvertData("currencies", symbols, converts)
}

//GetCurrencyPrice get price in string format
func (p *PricesInput) GetCurrencyPrice(symbol, convert string) string {
	return p.GetCurrency(symbol, convert)[0].Quotes[0].Price
}

//GetCurrencyPriceFloat64 get price formated to float64
func (p *PricesInput) GetCurrencyPriceFloat64(symbol, convert string) float64 {
	priceFloat, err := strconv.ParseFloat(p.GetCurrency(symbol, convert)[0].Quotes[0].Price, 64)
	if err != nil {
		log.Printf("Error on price parse: %s", err.Error())
	}
	return priceFloat
}

//GetCurrenciesQuotes get quotes objects with muiltple coins
func (p *PricesInput) GetCurrenciesQuotes(symbol, convert string) []Quote {
	return p.GetCurrencies(symbol, convert)[0].Quotes
}

//GetRequest general get request
func (p *PricesInput) GetRequest(url, token string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-token", token)
	resp, err := p.HTTPClient.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func setSecondsDuration(seconds int) time.Duration {
	return time.Second * time.Duration(seconds)
}
