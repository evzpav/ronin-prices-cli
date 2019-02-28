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

	"github.com/evzpav/ronin-prices-cli/prices"
)

func main() {

	p := prices.NewClient()

	currencyFullResp := p.GetCurrency("BTC", "TUSD")
	fmt.Printf("currency: %+v \n", currencyFullResp)

	ethBtcPrice := p.GetCurrencyPrice("ETH", "BTC")
    fmt.Printf("ETH/BTC price: %s \n", ethBtcPrice)
    
    ethBtcPrice := p.GetCurrencyPriceFloat64("ETH", "BTC")
	fmt.Printf("ETH/BTC price: %f \n", ethBtcPrice)

	currenciesFullResp := p.GetCurrencies("BTC,ETH", "BRL,TUSD")
	fmt.Printf("currencies: %+v \n", currenciesFullResp)

	quotesArray := p.GetCurrenciesQuotes("BTC,ETH", "BRL,TUSD")
	fmt.Printf("quotesArray: %+v \n", quotesArray)
}

```