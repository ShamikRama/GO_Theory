package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func change(p **Person) {
	*p = &Person{
		Name: "Alex",
		Age:  30,
	}
}

func change2(arr *[4]int) {
	arr[0] = 4
}

func main() {
	per := &Person{
		Name: "VVV",
		Age:  23,
	}
	change(&per)
	fmt.Println(per.Name, per.Age)
	arr := [...]int{1, 2, 3, 4}
	fmt.Println(arr)
	change2(&arr)
	fmt.Println(arr)
}
