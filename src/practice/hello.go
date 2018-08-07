package main
//Always Capitalize Variable names to be exported.Best practice
//gofmt <goscript> helps you fix style standard errors
//Cant use single quotes for strings

import("fmt")

func main(){
	i :=0
	for {
		fmt.Println(i)
		i +=2
		if i > 12{
			break
		}


	}
	fmt.Println("Stoped")
	//for _,Location := range s.Locations{
	//fmt.Println("\n%s",Location)

}

/*

import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func main(){
	resp, _  := http.Get("url")
	bytes,_ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	resp.Body.Close()
}
#####
//func contents should have tab alignment
//[5] or [4 5] <type> == array
//[] <type> ==slice
//mapper := make(map[String]float32)
//mapper["TestKey] = 100 ##Set values Just like Python
//delete(mapper,"TestKey")
//func test{}
//go test() causes concurrency
//sync.waitgroup
//waitgroup(1)
//go test() causes program to wait for other concurrent threads
//channels <- are to send over go routines and recieve processed values in return
//go routines are concurrent and non blocking but go channels sending and receiving values are blocking
//buffer to make channels return multiple value and be non blocking
//wg.wait()
//func test{
//defer any statement waits untill rest of the func is executed.Good for identifying job completion
//"panic" for exception conditions and "recover" to recover and continue
}



const usixteenbitmax float64 = 65535
const km_hr  = 1.60934

type car struct {
	gas uint64
	steering int16
	brake uint64
	top_speed float64


}
func (c car) kmh() float64 {
	return float64(c.gas) * (c.top_speed /usixteenbitmax)

}

func (c car) mmh() float64 {
	return float64(c.gas) * (c.top_speed /usixteenbitmax) /km_hr

}
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed = newspeed


}

func main(){
	a_car := car{65000,12561,0,225.0}
	fmt.Println(a_car.gas)
	a_car.new_top_speed(500)
	fmt.Println(a_car.mmh())
}


import ("fmt"
	"net/http"

)

func add(x , y float32) float32  {
	return x + y

}

func index_handle(resp http.ResponseWriter,req *http.Request)  {
	fmt.Fprintf(resp,"Hello Worlds!")

}

func about_handle(resp http.ResponseWriter,req *http.Request)  {
	fmt.Fprintf(resp,"This is go Test")

}

func main(){

	http.HandleFunc("/",index_handle)
	http.HandleFunc("/about",about_handle)
	http.ListenAndServe(":8000",nil)

	}

*/


