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
		APIEndpoint: "http://prices.stratum.bt",
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

//SetAPIEndpoint set new url
func (p *PricesInput) SetAPIEndpoint(url string) {
	p.APIEndpoint = url
}

//GetCurrencyConvertData reusable function
func (p *PricesInput) GetCurrencyConvertData(uri, symbol, convert string) (RoninPricesResp, error) {
	url := fmt.Sprintf("%s/%s?symbol=%s&convert=%s&source=%s", p.APIEndpoint, uri, symbol, convert, p.Source)
	body, err := GetRequest(p.HTTPClient, url, p.APIToken)
	if err != nil {
		log.Printf("Could not make request: %s", err.Error())
		return RoninPricesResp{}, err
	}
	var resp RoninPricesResp
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Printf("Error unmarshal: %s", err.Error())
		return RoninPricesResp{}, err
	}
	return resp, nil
}

//GetCurrency get currency full return
func (p *PricesInput) GetCurrency(symbol, convert string) (RoninPricesResp, error) {
	resp, err := p.GetCurrencyConvertData("currency", symbol, convert)
	if err != nil {
		return RoninPricesResp{}, err
	}
	return resp, nil
}

//GetCurrencies get currencies full return
func (p *PricesInput) GetCurrencies(symbols, converts string) (RoninPricesResp, error) {
	resp, err := p.GetCurrencyConvertData("currencies", symbols, converts)
	if err != nil {
		return RoninPricesResp{}, err
	}
	return resp, nil
}

//GetCurrencyPrice get price in string format
func (p *PricesInput) GetCurrencyPrice(symbol, convert string) (string, error) {
	resp, err := p.GetCurrency(symbol, convert)
	if err != nil {
		return "", err
	}
	return resp[0].Quotes[0].Price, nil
}

//GetCurrencyPriceFloat64 get price formated to float64
func (p *PricesInput) GetCurrencyPriceFloat64(symbol, convert string) (float64, error) {
	priceString, err := p.GetCurrencyPrice(symbol, convert)
	if err != nil {
		return float64(0), err
	}
	priceFloat, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		log.Printf("Error on price parse: %s", err.Error())
		return float64(0), err
	}
	return priceFloat, nil
}

//GetCurrenciesQuotes get quotes objects with muiltple coins
func (p *PricesInput) GetCurrenciesQuotes(symbol, convert string) ([]Quote, error) {
	quotes, err := p.GetCurrencies(symbol, convert)
	if err != nil {
		return []Quote{}, err
	}
	return quotes[0].Quotes, nil
}

//GetRequest general get request with x-token on header
func GetRequest(httpClient *http.Client, url, token string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-token", token)
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func setSecondsDuration(seconds int) time.Duration {
	return time.Second * time.Duration(seconds)
}
