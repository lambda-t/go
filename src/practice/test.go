package main

import (
	"math/rand"

	"fmt"
)

func main(){

	//Create Channels to receive call backs form goroutines

	c1 := make(chan int)
	done:= make(chan bool)

	//Spawn goroutines to add random numbers
		go func(done chan bool) {
			for i :=0;i <=99;i ++ {

				a := rand.Intn(99)
				b := rand.Intn(99)

				c := a + b
				//Send aggregate to channel c1
				c1 <- c


			}
			//Set flag to true to signal completion
			done <- true
		}(done)



    n:=0


	for i :=0;i <=99;i ++ {

		//receive values form go routines
		x := <-c1

		//Aggregate values
		n +=x


	}

	fmt.Println("Total is ", n)

	// Signal completion
	<- done




}


