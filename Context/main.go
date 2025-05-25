// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"time"
// 	"errors"
// )

// type resp struct {
// 	id int
// 	err error
// }

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	ctx := context.Background()
// 	ctx, _ = context.WithTimeout(ctx, time.Second*2)
// 	chanforResp := make(chan resp)
// 	go Call(ctx, chanforResp)
// 	res := <- chanforResp
// 	fmt.Println(res.id, res.err)

// }

// func Call(ctx context.Context, ch chan<- resp) {
// 	select{
// 	case <-ctx.Done():
// 		ch <- resp {
// 			id : 0,
// 			err: errors.New("timeout expired"),
// 		}
// 	case  <-time.After(time.Second * time.Duration(rand.Intn(5))):
// 		ch <- resp {
// 			id : rand.Intn(20),
// 			err : nil,
// 		}
// 	}
// }

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена по таймауту")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	makeRequest(ctx, "https://example.com")

	time.Sleep(6 * time.Second)
}

func makeRequest(ctx context.Context, url string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Запрос выполнен успешно")
}
