// // Структура в golang - это тип данных поле, занчение
// // Это не ссылочный тип данных
// // Когда мы неявно инициализируем структуру, например 
// type Humin struct{}
// h := Human{} // она пустая, но все равно имеет место в памяти, она всегда не nil
// // Проверка на пустоту структуры, можно сделать вот так:
// if h1 == (Human{}) {
//     fmt.Println("h1 is empty")
// }
// // Все остальные проверки h1 == 0 или h1 == nil приведут к ошибке компиляции

// // Это называется встраивание, аналог ООП 
// type Action struct {
//     val int
// }

// func (a Action) Get(){}

// type Human {
//     Action
// }

// chel := &Human{}

// chel.Get()

package main

import (
    "fmt"
    "sync"
)

func main(){
    //array := [4]int{2,3,4,5}
    fmt.Println(newarray(array))
}

func newarray(arr []int)[]int{
    mu := sync.WaitGroup{}
    newarray := make([]int, 0, len(arr))
    for _, val := range arr{
        go func(){
            mu.Lock()
            defer mu.Unlock()
            newarray := append(newarray, val * val)
        }()
    }
    return newarray
}

