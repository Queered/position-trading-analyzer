package main

import (
	"fmt"
	"github.com/adshao/go-binance"
)

func main() {

	binance := binance.NewClient("", "")


	symbol := "BTCUSDT" // CHANGE IT IF YOU WANT


	ticker, _, err := binance.Ticker.Price(symbol)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Current bid price: ", ticker.Bid)
	fmt.Println("Current ask price: ", ticker.Ask)


	order_book, _, err := binance.Depth.Get(symbol, binance.Depth.DefaultLimit)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Highest bid: ", order_book.Bids[0].Price)
	fmt.Println("Lowest ask: ", order_book.Asks[0].Price)
}
