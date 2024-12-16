// package main

// import (
// 	"fmt"
// 	"sync"
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
// 		time.Sleep(10 * time.Millisecond) // Имитация задержки чтения
// 	}
// }

// func main() {
// 	ch := make(chan int, 3) // Буферизированный канал с размером буфера 3
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go sender(ch)
// 	wg.Add(1)
// 	go receiver(ch)

// 	wg.Wait()
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

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func some(chj ...chan int) chan int {
// 	wg := sync.WaitGroup{}
// 	size := 0
// 	for _, c := range chj {
// 		size += len(c)
// 	}
// 	res := make(chan int, size)
// 	for _, cha := range chj {
// 		wg.Add(1)
// 		go func(cha chan int) {
// 			defer wg.Done()
// 			for val := range cha {
// 				res <- val
// 			}
// 		}(cha)
// 	}
// 	wg.Wait()
// 	close(res)
// 	return res
// }

// func main() {
// 	ch1 := make(chan int, 3)
// 	ch1 <- 1
// 	ch1 <- 2
// 	ch1 <- 3
// 	close(ch1)

// 	ch2 := make(chan int, 3)
// 	ch2 <- 4
// 	ch2 <- 5
// 	ch2 <- 6
// 	close(ch2)

// 	ch3 := make(chan int, 3)
// 	ch3 <- 7
// 	ch3 <- 8
// 	ch3 <- 9
// 	close(ch3)

// 	res := some(ch1, ch2, ch3)

// 	for val := range res {
// 		fmt.Println(val)
// 	}
// }

// func main() {
// 	ch := make(chan int, 7) // Буферизированный канал с емкостью 3

// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 7; i++ {
// 			ch <- i
// 			fmt.Println("Wrote value to channel:", i)
// 			//time.Sleep(500 * time.Millisecond) // Задержка для наглядности
// 		}
// 		//close(ch)
// 	}()
// 	//wg.Wait()
// 	//close(ch)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 7; i++ {
// 			fmt.Println("Read value from channel:", <-ch)
// 			//time.Sleep(1 * time.Second) // Задержка для наглядности
// 		}
// 		close(ch)
// 		fmt.Println("Channel closed")
// 	}()
// 	wg.Wait()
// 	fmt.Println("пиздец")

// 	for value := range ch {
// 		fmt.Println(value)
// 	}
// 	//time.Sleep(18 * time.Second)

// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	cnt := 500
// 	m := make(chan string, cnt)
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < cnt; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			m <- fmt.Sprintf("goroot %d", i)
// 		}(i)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(m)
// 	}()

// 	for val := range m {
// 		fmt.Println(val)
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	ch := make(chan int)
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			ch <- i
// 		}()
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()
// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// }

package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 2) // Буферизованный канал с размером буфера 2
	done := make(chan bool)
	var wg sync.WaitGroup

	// Добавляем одну горутину в WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			i := i
			select {
			case jobs <- i:
				fmt.Println("Отправлено", i)
			default:
				fmt.Println("Канал заполнен, пропускаем отправку")
			}
		}
		close(jobs) // Закрываем канал после отправки всех данных
	}()

	go func() {
		wg.Wait() // Ждем, пока все горутины завершат работу
		done <- true
	}()

	for j := range jobs {
		fmt.Println("Получено задание:", j)
	}

	<-done
	fmt.Println("Все задания получены")
}
