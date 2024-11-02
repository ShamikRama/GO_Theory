// Здесь будут задачи

// 1---------------------------------
// Дано: два неупорядоченных среза.
// а) a := []int{37, 5, 1, 2} и b := []int{6, 2, 4, 37}.
// б) a = []int{1, 1, 1} и b = []int{1, 1, 1, 1}.
// Верните их пересечение.

// func per(a, b []int) (res []int) {
// 	counter := map[int]int
// 	for _, elem := range a {
// 		if counter[elem] == 0 { // проверка существует ли элемент в мапе
// 			counter[elem] = 1
// 		} else {
// 			counter[elem] += 1
// 		}
// 	}
// 	for _, elem := range b {
// 		if counter[elem] > 0 { // проверка существует ли элемент в мапе
// 			counter[elem] -= 1
// 			res = append(res, elem)
// 		}
// 	}
// 	return res
// }

// -------------------------------

package main

import "fmt"

func main() {
	//var maps map[int]int
	maps := make(map[int]int)
	fmt.Println(maps == nil, len(maps))
	if _, ok := maps[0]; !ok {
		fmt.Println("map is empty")
	}
}

// --------------------------------

// package main

// import (
//     "fmt"
//     "runtime"
// )

// func main() {
//     var m runtime.MemStats
//     runtime.ReadMemStats(&m)
//     fmt.Printf("Memory usage before: %d bytes\n", m.Alloc)

//     s := make(map[int]int, 100)
//     for i := 0; i < 100; i++ {
//         s[i] = i
//     }

//     runtime.ReadMemStats(&m)
//     fmt.Printf("Memory usage after: %d bytes\n", m.Alloc)
// }
