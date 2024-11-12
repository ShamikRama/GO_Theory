// package main

// import (
// 	"fmt"
// 	//"time"
// 	"sync"
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

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func some(chj ...chan int) chan int {
// 	wg := sync.WaitGroup{}
// 	counter := 0
// 	for _, val := range chj {
// 		counter += len(val)
// 	}
// 	res := make(chan int, counter)
// 	for _, che := range chj {
// 		wg.Add(1)
// 		go func(che chan int) {
// 			defer wg.Done()
// 			for val := range che {
// 				res <- val
// 			}
// 		}(che)
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

// func main(){
// 	cnt := 500
// 	m := make(chan string, cnt)
// 	wg := sync.WaitGroup{}
// 	for i := 0 ; i < cnt; i++{
// 		wg.Add(1)
// 		go func (i int){
// 			defer wg.Done()
// 			m <- fmt.Sprintf("goroot %d", i)
// 		}(i)
// 	}
// 	//close(m)
// 	for i := 0; i < cnt; i++{
// 		wg.Add(1)
// 		go func(){
// 			defer wg.Done()
// 			Received(m)
// 		}()
// 	}
// 	//close(m)
// 	wg.Wait()
// }


// func Received(ch chan string){
// 	fmt.Println(<-ch)
// }

// func merge[T any](chans ...chan T) chan T{
// 	res := chan T
// 	wg := sync.WaitGroup{}
// 	for _, singlechan := range chans {
// 		wg.Add(1)
// 		singlechan := singlechan
// 		go func(){
// 			defer wg.Done()
// 			for val := range singlechan{
// 				res <- val
// 			}
// 		}()
// 	}
// 	go func(){
// 		wg.Wait()
// 		close(res)
// 	}
// 	return res
// }

// func main(){
// 	ch := make(chan int)
// 	wg := sync.WaitGroup{}
// 	for i:=0; i<20; i++{
// 		wg.Add(1)
// 		go func(){
// 			defer wg.Done()
// 			ch<-i
// 		}()
// 	}
// 	for i:=0; i<5; i++{
// 		wg.Add(1)
// 		go func(){
// 			defer wg.Done()
// 			fmt.Println(<-ch)
// 		}()
// 	}
// 	wg.Add(1)
// 	go func(){
// 		wg.Wait()
// 		close(ch)
// 	}()
// 	//time.Sleep(time.Second * 1)
// }

// package main

// import "fmt"

// func main() {
//     jobs := make(chan int, 5)
//     done := make(chan bool)
// 	wg:= sync.WaitGroup{}
// 	wg.Add(1)
//     go func() {
// 		defer wg.Done()
//         for {
//             j, more := <-jobs
//             if more {
//                 fmt.Println("received job", j)
//             } else {
//                 fmt.Println("received all jobs")
//                 done <- true
//                 return
//             }
//         }
//     }()
	
	
//     for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go func(){
// 			defer wg.Done()
//         jobs <- i
//         fmt.Println("sent job", i)
// 		fmt.Println("sent all jobs")
//     }()
// }
// 	wg.Add(1)
// 	go func(){
// 		defer wg.Done()
// 		wg.Wait()
// 		close(jobs)
// 	}()
//     // fmt.Println("sent all jobs")
	
//     //<-done
// }


// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	jobs := make(chan int, 2) // Буферизованный канал с размером буфера 2
// 	done := make(chan bool)
// 	var wg sync.WaitGroup

// 	// Добавляем одну горутину в WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 5; i++ {
// 			i := i
// 			//select {
// 			 jobs <- i
// 				fmt.Println("Отправлено", i)
// 			//default:
// 			//	fmt.Println("Канал заполнен, пропускаем отправку")
// 			}
// 		}()
// 		close(jobs) // Закрываем канал после отправки всех данных
	

// 	go func() {
// 		wg.Wait() // Ждем, пока все горутины завершат работу
// 		done <- true
// 	}()

// 	for j := range jobs {
// 		fmt.Println("Получено задание:", j)
// 	}

// 	<-done
// 	fmt.Println("Все задания получены")
// }

// import (
// 	"fmt"
// 	//"sync"
// 	"time"
// )

// func main() {
// 		ch := make(chan string)
// 		var res string
		

// 			ch <- "A"	// пишем
// 			close(ch)

	
// 	res = <-ch  // читаем
	
// 		time.Sleep(1*time.Second)
// 		fmt.Println(res)
// }