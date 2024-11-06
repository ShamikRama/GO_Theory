package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"errors"
)

type resp struct {
	id int
	err error
}

func main() {
	rand.Seed(time.Now().Unix())
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*2)
	chanforResp := make(chan resp)
	go Call(ctx, chanforResp)
	res := <- chanforResp
	fmt.Println(res.id, res.err)

}

func Call(ctx context.Context, ch chan<- resp) {
	select{
	case <-ctx.Done():
		ch <- resp {
			id : 0,
			err: errors.New("timeout expired"),
		}
	case  <-time.After(time.Second * time.Duration(rand.Intn(5))):
		ch <- resp {
			id : rand.Intn(20),
			err : nil,
		}
	}
}