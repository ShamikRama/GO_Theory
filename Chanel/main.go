// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sender(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i
// 		fmt.Println("Sent:", i)
// 	}
// 	close(ch)
// }

// func receiver(ch chan int) {
// 	for {
// 		val, ok := <-ch
// 		if !ok {
// 			fmt.Println("Channel closed")
// 			return
// 		}
// 		fmt.Println("Received:", val)
// 		time.Sleep(500 * time.Millisecond) // Имитация задержки чтения
// 	}
// }

// func main() {
// 	ch := make(chan int, 3) // Буферизированный канал с размером буфера 3

// 	go sender(ch)
// 	time.Sleep(40 * time.Second)
// 	go receiver(ch)

// 	// Даем время горутинам выполниться
// 	time.Sleep(5 * time.Second)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sender(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i
// 		fmt.Println("Sent:", i)
// 	}
// 	close(ch)
// }

// func receiver(ch chan int) {
// 	for {
// 		val, ok := <-ch
// 		if !ok {
// 			fmt.Println("Channel closed")
// 			return
// 		}
// 		fmt.Println("Received:", val)
// 		time.Sleep(500 * time.Millisecond) // Имитация задержки чтения
// 	}
// }

// func main() {
// 	ch := make(chan int) // Небуферизированный канал

// 	go sender(ch)
// 	go receiver(ch)

// 	// Даем время горутинам выполниться
// 	time.Sleep(10 * time.Second)
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	res := []int{}
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	time.Sleep(10 * time.Second)
	val := <-ch
	time.Sleep(10 * time.Second)
	close(ch)
	fmt.Println(val)

	for v := range ch {
		res = append(res, v)
	}
	fmt.Println(res)
}
