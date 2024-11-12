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



// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var sum int
// 	numbers := []int{1, 2, 3, 4, 5}
// 	var wg sync.WaitGroup

// 	for _, num := range numbers {
// 		wg.Add(1)
// 		go func(n int) {
// 			sum += n
// 			wg.Done()
// 		}(num)
// 	}

// 	wg.Wait()
// 	fmt.Println("Sum:", sum)
// }



package main

import (
    "fmt"
    "sync"
)


func Newarray(arr [4]int, new chan<-int){
	wg := sync.WaitGroup{}
	for _, val := range arr{
		wg.Add(1)
        go func(val int){
			defer wg.Done()
            new <- val
        }(val)
    }
	wg.Wait()
	close(new)
}

func main(){
    array := [4]int{2,3,4,5}
	newchannel := make(chan int, len(array))
	Newarray(array, newchannel)
	for val := range newchannel{
		fmt.Println(val)
	}
}

