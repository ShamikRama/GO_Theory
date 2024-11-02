// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func change(p *Person) {
// 	p = &Person{
// 		Name: "Alex",
// 		Age:  30,
// 	}
// }

// func main() {
// 	per := &Person{
// 		Name: "VVV",
// 		Age:  23,
// 	}
// 	change(per)
// 	fmt.Println(per.Name, per.Age)

// }
package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int
	numbers := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			sum += n
			wg.Done()
		}(num)
	}

	wg.Wait()
	fmt.Println("Sum:", sum)
}
