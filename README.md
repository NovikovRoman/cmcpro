# Coinmarketcap Pro API Client

Coinmarketcap Pro API Client written in Golang

Current version: `v1`

## Usage
```go
package main

import (
	"fmt"
	cmc "github.com/NovikovRoman/cmcpro"
	"log"
	"os"
	"time"
)

func main() {
	// true - https://pro-api.coinmarketcap.com/v1, false - https://sandbox-api.coinmarketcap.com/v1
	proxy := ""
	client, err := cmc.New(os.Getenv("API-KEY"), false, proxy, time.Duration(time.Second * 15))
	if err!= nil {
		log.Fatalf("create client %s", err.Error())
	}
	data, status, err := client.CryptocurrencyListingsLatest(
		10, 3, "name", cmc.SortAsc, cmc.NewConvertByCodes("RUB", "USD"), cmc.All)
	if err != nil {
		log.Fatalf("request %s", err.Error())
	}
	if status.ErrorMessage != "" {
		log.Fatalf("error %s", status.ErrorMessage)
	}
	fmt.Println(data)
}
```

See also tests

## Features
| Type           | Endpoint                               | Done |
|----------------|----------------------------------------|-------------|
| Cryptocurrency | /v1/cryptocurrency/info                | ✔ |
| Cryptocurrency | /v1/cryptocurrency/map                 | ✔ |
| Cryptocurrency | /v1/cryptocurrency/listings/latest     | ✔ |
| Cryptocurrency | /v1/cryptocurrency/listings/historical | ✔ |
| Cryptocurrency | /v1/cryptocurrency/market-pairs/latest | ✔ |
| Cryptocurrency | /v1/cryptocurrency/ohlcv/latest        | ✔ |
| Cryptocurrency | /v1/cryptocurrency/ohlcv/historical    | ✔ |
| Cryptocurrency | /v1/cryptocurrency/quotes/latest       | ✔ |
| Cryptocurrency | /v1/cryptocurrency/quotes/historical   | ✔ |
| Exchange       | /v1/exchange/info                      | ✔ |
| Exchange       | /v1/exchange/map                       | ✔ |
| Exchange       | /v1/exchange/listings/latest           | ✔ |
| Exchange       | /v1/exchange/listings/historical       | - |
| Exchange       | /v1/exchange/market-pairs/latest       | ✔ |
| Exchange       | /v1/exchange/quotes/latest             | ✔ |
| Exchange       | /v1/exchange/quotes/historical         | ✔ |
| Global Metrics | /v1/global-metrics/quotes/latest       | ✔ |
| Global Metrics | /v1/global-metrics/quotes/historical   | ✔ |
| Tools          | /v1/tools/price-conversion             | ✔ |

## Reference
[Coinmarketcap Pro v1](https://pro.coinmarketcap.com/api/v1)
