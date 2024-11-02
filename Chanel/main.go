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
	"sync"
)

func some(chj ...chan int) chan int {
	wg := sync.WaitGroup{}
	counter := 0
	for _, val := range chj {
		counter += len(val)
	}
	res := make(chan int, counter)
	for _, che := range chj {
		wg.Add(1)
		go func(che chan int) {
			defer wg.Done()
			for val := range che {
				res <- val
			}
		}(che)
	}
	wg.Wait()
	close(res)
	return res
}

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	ch2 := make(chan int, 3)
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	close(ch2)

	ch3 := make(chan int, 3)
	ch3 <- 7
	ch3 <- 8
	ch3 <- 9
	close(ch3)

	res := some(ch1, ch2, ch3)

	for val := range res {
		fmt.Println(val)
	}
}
