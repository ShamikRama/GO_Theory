package main

// func main() {
// 	counter := 20
//
// 	for i := 0; i <= counter; i++ {
//
// 		go func() {
//
// 			fmt.Println(i * i)
// 		}()
// 	}
//
// }

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
