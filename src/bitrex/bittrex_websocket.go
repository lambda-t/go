package main

import (
	"log"

	"github.com/carterjones/signalr"
	"github.com/carterjones/signalr/hubs"
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
)

// For more extensive use cases and capabilities, please see
// https://github.com/carterjones/bittrex.

func one(combo string) {
	// Prepare a SignalR client.
	c := signalr.New(
		"socket.bittrex.com",
		"1.5",
		"/signalr",
		`[{"name":"c2"}]`,
		nil,
	)

	// Set the user agent to one that looks like a browser.
	//c.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"

	// Send note to user about CloudFlare.
	//log.Println("Bypassing CloudFlare. This takes about 5 seconds.")


	//var m Message
	//json.Unmarshal(message, &m)

	//fmt.Printf("{EventType:%+v,EvenTime:%+v,Symbol:%+v,FirstUpdateId:%+v,FinalUpdateId:%+v," +
		//"Bids:%+v,Asks:%+v\n", m.EventType,m.EventTime,m.Symbol,m.FirstUpdateId,m.FinalUpdateId,m.Bids,m.Asks)


	/*

	msgHandler := func(msg signalr.Message) {
		var m Message
		json.Unmarshal(msg, &m)
	}
	*/
	//fmt.Printf("Exchange:Binance,s:%+v,E:%+v,U:%+v,u:%+v,b:u:%+v,a:u:%+v\n", m.S,m.E,  m.U, m.Z,m.B,m.A)


	// Define message and error handlers.
	msgHandler := func(msg signalr.Message) {

		//fmt.Println(msg.M)
		//var buf bytes.Buffer
		//enc := gob.NewEncoder(&buf)
		//enc.Encode(msg.M)
		//fmt.Print(buf.Bytes())
		//var value interface{}
		//var ok bool

		//jsonParsed,_:= gabs.ParseJSON(buf.Bytes())
		//fmt.Print(string(buf.Bytes()))
		//var f interface{}
		//json.Unmarshal(msg.M, &f)
		//json.Unmarshal(msg.M, &m)


	rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		fmt.Println(msg.M) // this gets captured

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		//fmt.Print(string(out[3]))
		re := regexp.MustCompile("Buys:(?s)(.*)Sells")
		sub1 := re.FindAllString(string(out),1)

        //v,_ := json.Marshal(string(out))
        fmt.Print(sub1)

		//var m Message
		//json.Unmarshal(out,&m)

		//fmt.Printf("{Buys:%+v", m)

		//fmt.Printf("Captured: %s", string(out)) // prints: Captured: Hello, playground

		//Set("Bittrex:" + combo,out)
		//var res,_ = Get("Bittrex:" + combo)
		//fmt.Print("res:",string(res))
		//var m = msg
		//var message := fmt.Sprintf(msg)

	/*
	buf := &bytes.Buffer{}
	msgHandler := func(msg signalr.Message) {
		binary.Write(buf,binary.LittleEndian,msg)
		fmt.Print(buf.ReadByte())
		//fmt.Println(msg)
	*/


		//Set("Bittrex:" + combo,msg)
	}



	panicIfErr := func(err error) {
		if err != nil {
			log.Panic(err)
		}
	}

	// Start the connection.
	err := c.Run(msgHandler, panicIfErr)
	panicIfErr(err)

	// Subscribe to the USDT-BTC feed.
	err = c.Send(hubs.ClientMsg{
		H: "corehub",
		M: "SubscribeToExchangeDeltas",
		A: []interface{}{combo},
		I: 1,
	})
	panicIfErr(err)

	// Wait indefinitely.
	select {}
}



func main(){

	args := os.Args[1:]
	fmt.Print(args)

	for i:= range args{
		fmt.Println(i)
		go one(args[i])
	}
	//go one("bnbbtc","addr1")
	//go one("ethbtc","addr2")
	//go one("bnbeth","addr3")
	fmt.Scanln()

}


//[{0 CoreHub updateExchangeState [map[Fills:[] MarketName:BTC-ETH Nounce:12211 Buys:[map[Type:1 Rate:0.06754827 Quantity:0] map[Type:0 Rate:0.0633 Quantity:0.08920906]] Sells:[map[Type:0 Rate:0.0731274 Quantity:6.83738251] map[Type:1 Rate:0.08025 Quantity:0]]]] <nil>}] 0    [] [] [] []}


//{d-4

//[{0 CoreHub updateExchangeState [map[Sells:[map[Rate:0.07905 Quantity:18.02419021 Type:2] map[Type:0 Rate:0.0795662 Quantity:6.28407558] map[Type:1 Rate:0.0800233 Quantity:0] map[Quantity:12.41193495 Type:0 Rate:0.08003037] map[Quantity:0 Type:1 Rate:0.08231763]] Fills:[map[Rate:0.07905 Quantity:0.54608121 TimeStamp:2018-06-08T06:37:47 OrderType:BUY]] MarketName:BTC-ETH Nounce:30345 Buys:[map[Type:0 Rate:0.07890008 Quantity:0.0942203] map[Type:1 Rate:0.07890002 Quantity:0] map[Type:0 Rate:0.078628 Quantity:2.58295] map[Quantity:0 Type:1 Rate:0.078627] map[Type:0 Rate:0.07841123 Quantity:26.4541] map[Type:1 Rate:0.07841121 Quantity:0]]]] <nil>}]

//[{0 CoreHub updateExchangeState [map[MarketName:BTC-ETH Nounce:25960 Buys:[map[Type:0 Rate:0.07525901 Quantity:1.552] map[Quantity:6.65283992 Type:0 Rate:0.07515587] map[Quantity:0 Type:1 Rate:0.07515523] map[Type:1 Rate:0.0695 Quantity:0]] Sells:[map[Type:1 Rate:0.0768936 Quantity:0] map[Type:0 Rate:0.0774549 Quantity:11.4013]] Fills:[]]] <nil>}]