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
func (c *PricesInput) SetPriceSource(source string) {
	c.Source = source
}

//ChangeHTTPClient set new http client if needed
func (c *PricesInput) ChangeHTTPClient(newHTTPClient *http.Client) {
	c.HTTPClient = newHTTPClient
}

//ChangeTimeout set new timeout in seconds
func (c *PricesInput) ChangeTimeout(newTimeout int) {
	c.HTTPClient.Timeout = setSecondsDuration(newTimeout)
}

func (c *PricesInput) GetCurrencyConvertData(uri, symbol, convert string) RoninPricesResp {
	url := fmt.Sprintf("%s/%s?symbol=%s&convert=%s&source=%s", "http://prices.stratum.bt", uri, symbol, convert, c.Source)
	body := c.GetRequest(url, c.APIToken)

	var resp RoninPricesResp
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Printf("Error unmarshal: %s", err.Error())
	}
	return resp
}

//GetCurrency get currency full return
func (c *PricesInput) GetCurrency(symbol, convert string) RoninPricesResp {
	return c.GetCurrencyConvertData("currency", symbol, convert)
}

//GetCurrencies get currencies full return
func (c *PricesInput) GetCurrencies(symbols, converts string) RoninPricesResp {
	return c.GetCurrencyConvertData("currencies", symbols, converts)
}

//GetCurrencyPrice get price in string format
func (c *PricesInput) GetCurrencyPrice(symbol, convert string) string {
	return c.GetCurrency(symbol, convert)[0].Quotes[0].Price
}

//GetCurrencyPriceFloat64 get price formated to float64
func (c *PricesInput) GetCurrencyPriceFloat64(symbol, convert string) float64 {
	priceFloat, err := strconv.ParseFloat(c.GetCurrency(symbol, convert)[0].Quotes[0].Price, 64)
	if err != nil {
		log.Printf("Error on price parse: %s", err.Error())
	}
	return priceFloat
}

//GetCurrenciesQuotes get quotes objects with muiltple coins
func (c *PricesInput) GetCurrenciesQuotes(symbol, convert string) []Quote {
	return c.GetCurrencies(symbol, convert)[0].Quotes
}

func (c *PricesInput) GetRequest(url, token string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-token", token)
	resp, err := c.HTTPClient.Do(req)

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
