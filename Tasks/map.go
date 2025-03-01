// package main

// type MyMap struct {
// 	mu   sync.RWMutex
// 	data map[string]string
// }

// func NewMap() *MyMap {
// 	data := make(map[string]string)
// 	return &MyMap{
// 		data: data,
// 	}
// }

// func (m *MyMap) GetOrCreate(key, value string) string {
// 	m.mu.RLock()
// 	val, ok := m.data[key]
// 	m.mu.RUnlock()

// 	if ok {
// 		return val
// 	}

// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	if val, ok := m.data[key]; ok {
// 		return val
// 	}

// 	m.data[key] = value
// 	return value
// }

// func main() {
// 	cm := NewMap()

// 	wg := sync.WaitGroup{}

// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		val := cm.GetOrCreate("key1", "value1")
// 		fmt.Println("go1: ", val)
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		val := cm.GetOrCreate("key1", "value2")
// 		fmt.Println("go2: ", val)
// 	}()

// 	wg.Wait()
// }

// package main

// import "fmt"

// // Надо удалить самые старые элементы в мапе
// type WordCounter struct {
// 	data    map[string]int
// 	limit   int
// 	counter []string
// }

// func NewWordCounter(limit int) *WordCounter {
// 	return &WordCounter{
// 		data:    make(map[string]int),
// 		limit:   limit,
// 		counter: make([]string, 0, limit),
// 	}
// }

// func (wc *WordCounter) CountWord(word string) {
// 	if _, ok := wc.data[word]; !ok {
// 		wc.counter = append(wc.counter, word)
// 	}
// 	wc.data[word]++

// 	if len(wc.data) > wc.limit {
// 		delete(wc.data, wc.counter[0])
// 		wc.counter = wc.counter[1:]
// 	}
// }

// func main() {
// 	wc := NewWordCounter(4)

// 	words := []string{"bmw", "merc", "yaguar", "rangerover", "buggatti"}

// 	for _, val := range words {
// 		wc.CountWord(val)
// 	}

// 	fmt.Println(wc.data)
// }

// package main

// import "fmt"

// // у нас есть мапа в которой ключ это стринг а занчение это массив строк
// // мы хотим вставить данные(новый слайс) в мапу по ключу если у нас есть уже
// // в занчении мапы в слайсе это слово то мы его не вставляем

// func mergeToMap(data map[string][]string, key string, newValue []string) {
// 	uMap := make(map[string]struct{})
// 	for _, val := range data[key] {
// 		uMap[val] = struct{}{}
// 	}

// 	for _, val := range newValue {
// 		if _, ok := uMap[val]; !ok {
// 			data[key] = append(data[key], val)
// 			uMap[val] = struct{}{} // здесь newValue могут быть дубликаты поэтому я сделал так
// 		}
// 	}
// }

// func main() {
// 	oldMap := map[string][]string{
// 		"group1": {"apple", "banana"},
// 		"group2": {"carrot"},
// 	}
// 	fmt.Println(oldMap)

// 	newValues := []string{"banana", "chery", "banana", "chery"}
// 	fmt.Println(newValues)

// 	mergeToMap(oldMap, "group1", newValues)
// 	fmt.Println(oldMap)
// }

package main

import (
	"fmt"
	"math"
)

// На Авито размещено множество товаров, каждый из которых представлен числом.
// У каждого покупателя есть потребность в товаре, также выраженная числом.
// Если точного товара нет, покупатель выбирает ближайший по значению товар,
// что вызывает неудовлетворённость, равную разнице между его потребностью и купленным товаром.
// Количество каждого товара не ограничено, и один товар могут купить несколько покупателей.
// Рассчитайте суммарную неудовлетворённость всех покупателей.

// Нужно написать функцию, которая примет на вход два массива:
// массив товаров и массив потребностей покупателей, вычислит сумму неудовлетворённостей всех покупателей и вернет результат в виде числа.

// __________________________________
// ПРИМЕР
// Ввод
// goods = [8, 3, 5]
// buyerNeeds = [5, 14, 12, 44, 55] // 6 + 4 + 36 + 47 = 93

// Вывод
// res = 1 # первый покупатель покупает товар 5 и его неудовлетворённость = 0, второй также покупает товар 5 и его неудовлетворённость = 6-5 = 1

// needs - product = unhappy
// 0 + 6 + 4 + 36 + 47 = 93

func FindTarget(gooods []int, target int) int {
	low := 0
	high := len(gooods) - 1

	for low <= high {
		mid := (high - low) / 2
		if gooods[mid] == target {
			return target
		} else if gooods[mid] < target {
			low = mid + 1
		} else {
			low = mid - 1
		}
	}

	if math.Abs(float64(target)-float64(gooods[low])) < math.Abs(float64(target)-float64(gooods[high])) {
		return gooods[low]
	}
	return gooods[high]

}

func FindGood(goods []int, needs []int) int {
	sum := 0
	for _, val := range needs {
		closest := FindTarget(goods, val)
		sum += int(math.Abs(float64(val)) - float64(closest))
	}
	return sum
}

func main() {
	goods := []int{8, 3, 5}
	buyerNeeds := []int{5, 14, 12, 44, 55}
	fmt.Println(FindGood(goods, buyerNeeds))
}
