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

var netClient = &http.Client{
	Timeout: time.Second * time.Duration(10),
}

func NewClient(token string) *PricesInput {
	return &PricesInput{
		Source:   "coinmarketcap",
		APIToken: token,
	}
}

func (c *PricesInput) SetPriceSource(source string) *PricesInput {
	c.Source = source
	return c
}

func (c *PricesInput) getCurrencyConvertData(uri, symbol, convert string) RoninPricesResp {
	url := fmt.Sprintf("%s/%s?symbol=%s&convert=%s&source=%s", "http://prices.stratum.bt", uri, symbol, convert, c.Source)
	body := getRequest(url, c.APIToken)

	var resp RoninPricesResp
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Printf("Error unmarshal: %s", err.Error())
	}
	return resp
}

func (c *PricesInput) GetCurrency(symbol, convert string) RoninPricesResp {
	return c.getCurrencyConvertData("currency", symbol, convert)
}

func (c *PricesInput) GetCurrencies(symbols, converts string) RoninPricesResp {
	return c.getCurrencyConvertData("currencies", symbols, converts)
}

func (c *PricesInput) GetCurrencyPrice(symbol, convert string) string {
	return c.GetCurrency(symbol, convert)[0].Quotes[0].Price
}

func (c *PricesInput) GetCurrencyPriceFloat64(symbol, convert string) float64 {
	priceFloat, err := strconv.ParseFloat(c.GetCurrency(symbol, convert)[0].Quotes[0].Price, 64)
	if err != nil {
		log.Printf("Error on price parse: %s", err.Error())
	}
	return priceFloat
}

func (c *PricesInput) GetCurrenciesQuotes(symbol, convert string) []Quote {
	return c.GetCurrencies(symbol, convert)[0].Quotes
}

func getRequest(url, token string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-token", token)
	resp, err := netClient.Do(req)

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
