package main

import "fmt"

func main() {
	s := "shalom😃"
	for _, char := range s {
		fmt.Println(string(char))
	}
	fmt.Println(len([]rune(s)))
}
