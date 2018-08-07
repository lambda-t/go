package main


import (
"github.com/toorop/go-bittrex"

	"fmt"
)

func main() {
	bt := bittrex.New("", "")
	ch := make(chan bittrex.ExchangeState)
	go func() {
		bt.SubscribeExchangeUpdate("USDT-BTC", ch, nil)

	}()

	//fmt.Println(len(ch), cap(ch))

	fmt.Println(<-ch)

	select {
	}

}
