// Здесь будут задачи

// 1-----------------------

// Задачи на понимание особенностей работы append и reslicing

// func foo(a []int) {
// 	a = append(a, 7)
// 	a[1] = 7
// }

// func bar(a *[]int) {
// 	*a = append(*a, 7)
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Printf("a[1]=%d\n", a[1]) // 2 +

// 	b := a[1:3]                      // [2 , 3]
// 	b[0] = 10                        // [10 , 3]
// 	fmt.Printf("1. a[1]=%d\n", a[1]) // 10 +
// 	// a [1,10,3,4,5,6]
// 	b = append(b, a...)              // b [10 , 3 , 1 ,10 , 3 , 4 ,5, 6]
// 	b[0] = 100                       // [100 , 3 , 1 ,10 , 3 , 4 ,5, 6]
// 	fmt.Printf("2. a[1]=%d\n", a[1]) // 10 +

// 	foo(a)
// 	fmt.Printf("3. a[1]=%d\n", a[1]) // 10 +

// 	bar(&a)
// 	fmt.Printf("4. a=%v\n", a) // [1,10,3,4,5,6,7] +

// }

//package main
//
//import "fmt"
//
//func main() {
//	a := [2]int{0, 0}
//
//	if a == nil {
//		fmt.Println("true")
//	} else {
//		fmt.Println("false")
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	b := new(bool)
//	modify(b)
//	fmt.Println(b)
//}
//
//func modify(b *bool) {
//	b = nil
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	var wg sync.WaitGroup
//	counter := 0
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			counter++
//		}()
//	}
//	wg.Wait()
//	fmt.Println(counter)
//}
