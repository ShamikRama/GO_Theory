package main

import "fmt"

// func main() {
// 	a1 := make([]int, 0, 10)
// 	a1 = append(a1, []int{1, 2, 3, 4, 5}...)
// 	a2 := append(a1, 6)
// 	a3 := append(a1, 7)
// 	a4 := append(a1, 45, 56, 67, 67)

// 	fmt.Println(a1, a2, a3, a4)
// }

// Надо удалить самые старые элементы в мапе

type WordCounter struct {
	data    map[string]int
	limit   int
	counter []string
}

func NewWordCounter(limit int) *WordCounter {
	return &WordCounter{
		data:    make(map[string]int),
		limit:   limit,
		counter: make([]string, 0, limit),
	}
}

func (wc *WordCounter) CountWord(word string) {
	if _, ok := wc.data[word]; !ok {
		wc.counter = append(wc.counter, word)
	}
	wc.data[word]++

	if len(wc.data) > wc.limit {
		delete(wc.data, wc.counter[0])
		wc.counter = wc.counter[1:]
	}
}

func main() {
	wc := NewWordCounter(4)

	words := []string{"bmw", "merc", "yaguar", "rangerover", "buggatti"}

	for _, val := range words {
		wc.CountWord(val)
	}

	fmt.Println(wc.data)
}
