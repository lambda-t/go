package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"net/url"
	"github.com/gorilla/websocket"
	"fmt"
	"time"
	"encoding/json"
)

func one(combo string, flag_name string)(){

	var addr = flag.String(flag_name, "api.hitbtc.com", "http service address")
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "api/2/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)



	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	func()string {
		type Params  map[string]string
		type Req struct {
			Method string `json:"method"`
			Params `json:"params"`
			Id int `json:"id"`
		}



		r:= Req{"subscribeOrderbook",Params{"symbol":combo},123}

		c.WriteJSON(r)



		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return "error"
			}
			//return string(message)

			type Message struct {
				Jsonrpc string `json:"_"`
				Method  string `json:"method"`
				Params  struct {
					Asks [] struct {
						Price string `json:"price"`
						Size  string `json:"size"`
					} `json:"ask"`
					Bids [] struct {
						Price string `json:"price"`
						Size  string `json:"size"`
					} `json:"bid"`
					Symbol   string `json:"symbol"`
					UpdateId  int    `json:"sequence"`
				} `json:"params"`
			}



			var m Message
			json.Unmarshal(message, &m)


			res := map[string]interface{}{"EventType":m.Method,"Symbol":m.Params.Symbol,"UpdateId":m.Params.UpdateId,
				"Bids":m.Params.Bids,"Asks":m.Params.Asks}
			res1,_ := json.Marshal(res)

			//fmt.Println(string(res1))
			Set("HiBtc:" + combo,res1)


			//fmt.Print("recived:", string(message) + "\n")

			//type ParseMessage struct {
				//exchange string
				//Message_Binance string
			//}
			//type ConstructStream struct{
				//Stream ParseMessage


			//}



			//var m Message
			//json.Unmarshal(message, &m)


			/*
			data := []byte`{
				"EventType":m.
				EventTime
				Symbol
				FirstUpdateId
				FinalUpdateId
				Bids
				Asks
			}
			*/


			//fmt.Print(m)

			//{Binance" +combo + ":

			//fmt.Printf("{EventType:%+v,EvenTime:%+v,Symbol:%+v,FirstUpdateId:%+v,FinalUpdateId:%+v," +
			//"Bids:%+v,Asks:%+v\n", m.EventType,m.EventTime,m.Symbol,m.FirstUpdateId,m.FinalUpdateId,m.Bids,m.Asks)

			//fmt.Printf("Exchange:Binance,EventType:%+v,EvenTime:%+v,Symbol:%+v,FirstUpdateId:%+v,FinalUpdateId:%+v," +
				//"Bids:%+v,Asks:%+v\n", m.EventType,m.EventTime,m.Symbol,m.FirstUpdateId,m.FinalUpdateId,m.Bids,m.Asks)

			//v, _ := jason.NewObjectFromBytes([]byte(message))

			//fmt.Printf(v.GetString("U"))



			//d,_ := json.Marshal({m.EventType,m.EventTime})


			//Set("Binance:" + combo,message)
			//var r,_ = Get("Binance:" + combo)
			//fmt.Print("res:",string(message))

			//var r,_ = Get("Binance")
			//fmt.Print(string(message) + "\n")
			//"%+v\n",

			//data := &ConstructStream{Stream:ParseMessage{Exchange:"Binance",Message:string(message)}}
			//data  := &ParseMessage{Message_Binance:string(message)}
			//r,_ := json.Marshal(data)
			//Set("Binance:" + ,message)

			//fmt.Println(string(message))
			//fmt.Println(string(data))



			//Set("Binance",message)
			//data := []byte(message)
			//json.Marshal(CombinedStream(data))
			//mapD := map[string]byte {"Binance": json.Unmarshal(message, &m)}
			//mapB, _ := json.Marshal(mapD)
			//fmt.Println(string(mapB))
			//d := map[string][]byte{"Binance":message}
			//mapB, _ := json.Marshal(d)


			//Set("Stream",mapB)

			//E := []byte("Exchange:Binance,s:%+v,E:%+v,U:%+v,u:%+v,b:u:%+v,a:u:%+v\n", m.S,m.E,  m.U, m.Z,m.B,m.A)
			//Set("E1",E)
			//var r,_ = Get("Binance")
			//fmt.Print("%+v\n",string(r))
			//fmt.Printf("Exchange:Binance,s:%+v,E:%+v,U:%+v,u:%+v,b:u:%+v,a:u:%+v\n", m.S,m.E,  m.U, m.Z,m.B,m.A)

			//log.Printf("recv: %s", message)
		}
	}()



	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return


		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			/*
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					return
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
			*/
			return
		}
	}


}


func main(){

	args := os.Args[1:]
	fmt.Print(args)

	for i:= range args{
		fmt.Println(i)
		go one(args[i],"addr" + string(i+1))
	}
	//go one("bnbbtc","addr1")
	//go one("ethbtc","addr2")
	//go one("bnbeth","addr3")
	fmt.Scanln()

}









// {"e":"depthUpdate","E":1526186677776,"s":"BNBBTC","U":71377818,"u":71377822,"b":[["0.00152750","84.28000000",[]]],"a":[["0.00152970","0.00000000",[]],["0.00153160","82.40000000",[]]]}

//go run hibtc_websocket.go redis_init.go ETHBTC

//go run  redis_init.go ETHBTC


//docker run --net=host  -it --entrypoint "bash" binance_web_socket_receiver
//docker run --net=host  -it --name binance binance_web_socket_receiver
//docker run -it binance_web_socket_receiver bnbbtc --link freecache:freecache binance_web_socket_receiver
//docker build -t binance_web_socket_receiver .
//docker run -it --link freecache:freecache --name binance binance_web_socket_receiver bnbbtc

//docker run -it -p 6379 --link redis1:redis --name client2 redis sh

//recived:{"jsonrpc":"2.0","method":"snapshotOrderbook",
// "params":80130","size":"0{"ask":[{"price":"0.080126","size":"0.103"},{"price":"0.080129","size":"0.409"},
//// {"price":"0.0.001"},{"price":"0.080133","size":"0.001"},{"price":"0.080138","size":"0.680"}]
// ,"symbol":"ETHBTC","sequence":860363}}


//go run hibtc_websocket.go ETHBTC

