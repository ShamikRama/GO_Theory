package main

import (
	"fmt"
	//"rand"
	"context"
	"time"
)

func main(){
	rand.Seed(time.Now().Unix())
	ctx := context.Backroud()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	chanforResp := make(chan int)
	go Call(ChanforResp)

	select {
	case <- ctx.Done():
		fmt.Println("timeout")
	case result := <-chanforResp:
		fmt.Ptintln(result)
	}
	cancel()
}

func ChanforResp(ch chan<- int){
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
}