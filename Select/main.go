// package main

// import "fmt"

// func main() {
//     messages := make(chan string)
//     signals := make(chan bool)

//     select {
//     case msg := <-messages:
//         fmt.Println("received message", msg)
//     default:
//         fmt.Println("no message received")
//     }

//     msg := "hi"
//     select {
//     case messages <- msg:
//         fmt.Println("sent message", msg)
//     default:
//         fmt.Println("no message sent")
//     }

//     select {
//     case msg := <-messages:
//         fmt.Println("received message", msg)
//     case sig := <-signals:
//         fmt.Println("received signal", sig)
//     default:
//         fmt.Println("no activity")
//     }
// }

package main

//func main() {
//
//	c1 := make(chan string)
//	go func() {
//		time.Sleep(2 * time.Second)
//		c1 <- "result 1"
//	}()
//
//	select {
//	case res := <-c1:
//		fmt.Println(res)
//	case <-time.After(1 * time.Second):
//		fmt.Println("timeout 1")
//	}
//
//	c2 := make(chan string)
//	wg := sync.WaitGroup{}
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		time.Sleep(2 * time.Second)
//		c2 <- "result 2"
//	}()
//	//wg.Wait()
//	select {
//	case res := <-c2:
//		fmt.Println(res)
//	case <-time.After(3 * time.Second):
//		fmt.Println("timeout 2")
//	}
//}
