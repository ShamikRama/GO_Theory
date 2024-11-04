package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	counter := 5
	wg := sync.WaitGroup{}
	wg.Add(counter)

	for i := 0; i < counter; i++ {

		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

// func main() {
// 	writes := 1000
// 	var storage map[int]int
// 	wg := sync.WaitGroup{}
// 	wg.Add(writes)
// 	for i := 0; i < writes; i++ {
// 		i := i
// 		go func() {
// 			defer wg.Done()
// 			storage[i] = i
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(storage)
// }
