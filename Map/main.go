// Здесь будут задачи

// 1---------------------------------
// Дано: два неупорядоченных среза.
// а) a := []int{37, 5, 1, 2} и b := []int{6, 2, 4, 37}.
// б) a = []int{1, 1, 1} и b = []int{1, 1, 1, 1}.
// Верните их пересечение.

// func per(a, b []int) (res []int) {
// 	counter := map[int]int
// 	for _, elem := range a {
// 		if counter[elem] == 0 { // проверка не существует ли элемент в мапе
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

// package main

// import "fmt"

// func main() {
// 	var maps map[int]int
// 	//maps := make(map[int]int)
// 	fmt.Println(maps == nil, len(maps))
// 	if _, ok := maps[0]; !ok {
// 		fmt.Println("map is empty")
// 	}
// }

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

// package main

// import "fmt"

// type account struct {
// 	balance int
// }

// func main(){
// m:=make(map[string]*account)
// m["user1"]=&account{
// 	balance:100,
// }
// m["user1"].balance += 50
// 	//Написать что-то, чтобы увеличить баланс на 50
// fmt.Println(m["user1"].balance)

// }

// package main

// import "fmt"

// // дубликат
// func dub(obj []string) []string {
// 	res := make([]string, 0, len(obj))

// 	counter := make(map[string]int)

// 	for _, val := range obj {
// 		counter[val]++
// 		if counter[val] == 2 {
// 			res = append(res, val)
// 		}
// 	}

// 	return res
// }

// // недубликат
// func reversedub(obj []string) {
// 	counter := make(map[string]int)

// 	for _, val := range obj {
// 		counter[val]++
// 	}

// 	for key, val := range counter {
// 		if val == 1 {
// 			fmt.Printf("%s ", key)
// 		}
// 	}
// 	fmt.Println()
// }

// // пересечение из двух слайсов
// func intersection(a, b []string) []string {
// 	counter := make(map[string]bool)
// 	res := make([]string, 0)
// 	for _, elem := range a {
// 		counter[elem] = true
// 	}

// 	for _, elem := range b {
// 		if counter[elem] {
// 			res = append(res, elem)
// 		}
// 	}
// 	return res
// }

// // вывести све элементы только в одном экземпляре
// func oneobj(obj []string) {
// 	res := make(map[string]struct{})

// 	for _, val := range obj {
// 		res[val] = struct{}{}
// 	}

// 	fmt.Println(res)
// }

// func main() {
// 	arr := []string{"a", "bb", "a", "bb", "c", "d"}
// 	arr1 := []string{"v", "a", "h", "g"}
// 	arr2 := []string{"v", "a", "f", "k"}
// 	fmt.Println(dub(arr))
// 	reversedub(arr)
// 	oneobj(arr)
// 	fmt.Println(intersection(arr1, arr2))

// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type messageStruct struct {
// 	Message string
// }

// type DarkKernel struct {
// 	Messages   map[int]messageStruct
// 	MessageId  int
// 	IsShutDown bool
// 	Mu         *sync.RWMutex
// }

// type DarkInterface interface {
// 	shutdown()
// 	send(msg string) (int, error)
// 	recive(msgId int) (string, error)
// }

// func New() *DarkKernel {
// 	return &DarkKernel{
// 		Messages: make(map[int]messageStruct),
// 		Mu:       &sync.RWMutex{},
// 	}
// }

// func (d *DarkKernel) shutdown() {
// 	defer d.Mu.Unlock()
// 	d.Mu.Lock()
// 	d.IsShutDown = true
// }

// func (d *DarkKernel) send(msg string) (int, error) {
// 	if d.IsShutDown {
// 		panic("Fuck")
// 	}

// 	defer d.Mu.Unlock()
// 	d.Mu.Lock()

// 	d.MessageId++
// 	d.Messages[d.MessageId] = messageStruct{
// 		Message: msg,
// 	}

// 	return d.MessageId, nil
// }

// func (d *DarkKernel) recive(msgId int) (string, error) {
// 	defer d.Mu.Unlock()
// 	d.Mu.Lock()

// 	msg, ok := d.Messages[msgId]
// 	if !ok {
// 		return "", fmt.Errorf("ничего нет тут")
// 	}

// 	return msg.Message, nil

// }

package main

import "fmt"

// Поиск суммы в массиве
// Напиши функцию, принимающую массив из положительных неупорядоченных чисел
// первым аргументом и положительное число вторым аргументом.
// Функция должна возвращать true, если в массиве есть 2 числа,
// которые в сумме дают 2-й аргумент.

func foo(arr []int, n int) bool {
	seen := make(map[int]bool)

	for _, num := range arr {
		target := num - n
		if seen[target] {
			return true
		}

		seen[num] = true
	}

	return false

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo(arr, 4))
	fmt.Println(foo(arr, 12))
}
