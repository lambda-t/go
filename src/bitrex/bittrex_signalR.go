package main

import (
	"github.com/thebotguys/signalr"
	"encoding/json"
	"fmt"
)


func main(){
	client := signalr.NewWebsocketClient()

	client.OnClientMethod = func(hub, method string, arguments []json.RawMessage) {
		fmt.Println("Message Received: ")
		fmt.Println("HUB: ", hub)
		fmt.Println("METHOD: ", method)
		fmt.Println("ARGUMENTS: ", arguments)
	}
	client.OnMessageError = func (err error) {
		fmt.Println("ERROR OCCURRED: ", err)
	}

	c := client.Connect("https", "socket.bittrex.com", []string{"corehub"})
	fmt.Println(c)


	e,err := client.CallHub("corehub", "GET","SubscribeExchangeUpdate", "USDT-BTC")
	//e,_ := client.CallHub("corehub", "GET", "params", 1, 1.4, "every type is accepted")


	fmt.Println(e,err)
}
