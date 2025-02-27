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

package main

import "fmt"

// у нас есть мапа в которой ключ это стринг а занчение это массив строк
// мы хотим вставить данные(новый слайс) в мапу по ключу если у нас есть уже
// в занчении мапы в слайсе это слово то мы его не вставляем

func mergeToMap(data map[string][]string, key string, newValue []string) {
	uMap := make(map[string]struct{})
	for _, val := range data[key] {
		uMap[val] = struct{}{}
	}

	for _, val := range newValue {
		if _, ok := uMap[val]; !ok {
			data[key] = append(data[key], val)
			uMap[val] = struct{}{} // это вроде необязательно
		}
	}

}

func main() {
	oldMap := map[string][]string{
		"group1": {"apple", "banana"},
		"group2": {"carrot"},
	}
	fmt.Println(oldMap)

	newValues := []string{"banana", "chery"}
	fmt.Println(newValues)

	mergeToMap(oldMap, "group1", newValues)
	fmt.Println(oldMap)
}
