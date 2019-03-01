# ronin-prices-cli

## How to install:
``` bash
go get github.com/evzpav/ronin-prices-cli
```

## Example:

```go
package main

import (
	"fmt"
	"os"

	"github.com/evzpav/ronin-prices-cli/prices"
)

func main() {

	p := prices.NewClient(os.Getenv("TOKEN"))

	currencyFullResp, _ := p.GetCurrency("BTC", "TUSD")
	fmt.Printf("currency: %+v \n", currencyFullResp)

	ethBtcPrice, _ := p.GetCurrencyPrice("ETH", "BTC")
	fmt.Printf("ETH/BTC price: %s \n", ethBtcPrice)

	ethBtcPriceFloat, _ := p.GetCurrencyPriceFloat64("ETH", "BTC")
	fmt.Printf("ETH/BTC price: %.8f \n", ethBtcPriceFloat)

	currenciesFullResp, _ := p.GetCurrencies("BTC,ETH", "BRL,TUSD")
	fmt.Printf("currencies: %+v \n", currenciesFullResp)

	quotesArray, _ := p.GetCurrenciesQuotes("BTC,ETH", "BRL,TUSD")
	fmt.Printf("quotesArray: %+v \n", quotesArray)
}

```