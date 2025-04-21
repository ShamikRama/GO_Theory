package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	wg := sync.WaitGroup{}
	//mu := sync.Mutex{}
	wg.Add(12)

	for i := 0; i < 12; i++ {

		go func() {
			defer wg.Done()
			//mu.Lock()
			counter++
			//mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
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
