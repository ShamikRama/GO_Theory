package main

import "fmt"

func foo(a [5]int) {
	a[3] = 10
}

func bar(a *[5]int) {
	a[3] = 10
}

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	fmt.Printf("%#v\n", a) // []int{1, 2, 3, 4, 5}

	foo(a)
	fmt.Printf("%#v\n", a) // []int{1, 2, 3, 10, 5}

	bar(&a)
	fmt.Printf("%#v\n", a) // []int{1, 2, 3, 10, 5}
}
