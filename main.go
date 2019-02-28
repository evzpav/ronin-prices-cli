package main

import (
	"fmt"

	"github.com/evzpav/ronin-prices-cli/prices"
)

func main() {

	p := prices.NewClient("token")

	currencyFullResp := p.GetCurrency("BTC", "TUSD")
	fmt.Printf("currency: %+v \n", currencyFullResp)

	ethBtcPrice := p.GetCurrencyPrice("ETH", "BTC")
	fmt.Printf("ETH/BTC price: %s \n", ethBtcPrice)

	ethBtcPriceFloat := p.GetCurrencyPriceFloat64("ETH", "BTC")
	fmt.Printf("ETH/BTC price: %.8f \n", ethBtcPriceFloat)

	currenciesFullResp := p.GetCurrencies("BTC,ETH", "BRL,TUSD")
	fmt.Printf("currencies: %+v \n", currenciesFullResp)

	quotesArray := p.GetCurrenciesQuotes("BTC,ETH", "BRL,TUSD")
	fmt.Printf("quotesArray: %+v \n", quotesArray)
}
