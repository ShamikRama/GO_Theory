//package main

//=================================================================================================
//===========================     Слайсы, Мапы   ==================================================
//=================================================================================================

//func main() {
//	var foo []int
//	var bar []int
//
//	foo = append(foo, 1)  // 1 1 [1]
//	foo = append(foo, 2)  // 2 2 [1 2]
//	foo = append(foo, 3)  // 3 4 [1 2 3]
//	bar = append(foo, 4)  // [1 2 3 4]
//	foo = append(foo, 5)  // [1 2 3 5]
//	fmt.Println(foo, bar) // [1 2 3 5] [1 2 3 5]
//
//}

//func main() {
//	c := []string{"A", "B", "D", "E"}
//	b := c[1:2]
//	fmt.Println(len(b), cap(b)) // 1 3
//	b = append(b, "TT")         // [B TT] 2 3
//	fmt.Println(c)              // [A B TT E]
//	fmt.Println(b)              // [B TT]
//}

//func main() {
//	m := make(map[string]int)
//	for _, word := range []string{"hello", "world", "from", "the",
//		"best", "language", "in", "the", "world"} {
//		m[word]++
//	}
//	for k, v := range m {
//		fmt.Println(k, v)
//	}
//}

//func main() {
//	mutate := func(a []int) {
//		a[0] = 0
//		a = append(a, 1)
//		fmt.Println(a) // 0 2 3 4 1
//	}
//	a := []int{1, 2, 3, 4}
//	mutate(a)
//	fmt.Println(a) // 0 2 3 4
//
//}

//func main() {
//	sl := []int{1, 2, 3, 5}
//	mod(sl)
//	fmt.Println(sl) // 5 5 5 5
//}
//
//func mod(a []int) {
//	for i := range a {
//		a[i] = 5
//	}
//	fmt.Println(a) // 5 5 5 5
//}

//func main() {
//	sl := make([]int, 4, 8) // 0 0 0 0 ...
//	sl[0] = 1               // 1 0 0 0
//	sl[1] = 2               // 1 2 0 0
//	sl[2] = 3               // 1 2 3 0
//	sl[3] = 5               // 1 2 3 5
//	mod(sl)
//	fmt.Println(sl) // 5 5 5 5
//}
//
//func mod(a []int) {
//	for i := range a {
//		a[i] = 5
//	}
//	fmt.Println(a) // 5 5 5 5
//}

//func main() {
//	sl := []int{1, 2, 3, 4, 5}
//	mod(sl)
//	fmt.Println(sl) // 1 2 3 4 5
//}
//
//func mod(a []int) {
//	a = append(a, 125)
//	for i := range a {
//		a[i] = 5
//	}
//	fmt.Println(a) // [5 5 5 5 5 5]
//}

//func main() {
//	s := make([]int, 3, 8) // [0 0 0 ...] cap 8 len 3
//	m := make(map[int]int, 8)
//
//	a(s)
//	fmt.Println(s[3]) // panic
//
//	b(m)
//	fmt.Println(m[3]) // 33
//}
//
//func a(s []int) {
//	s = append(s, 37)
//}
//
//func b(m map[int]int) {
//	m[3] = 33
//}

//func main() {
//	a := []int{1, 2}  // len 2 cap 2
//	a = append(a, 3)  // len 3 cap 4 [1 2 3]
//	b := append(a, 4) // [1 2 3 4]
//	c := append(a, 5) // [1 2 3 5]
//
//	fmt.Println(b) // [1 2 3 5]
//	fmt.Println(c) // [1 2 3 5]
//}

//func main() {
//	a := []int{1, 2}  // len 2 cap 2
//	a = append(a, 3)  // [1 2 3] len 3 cap 4
//	a = append(a, 7)  // [1 2 3 7] len 4 cap 4
//	b := append(a, 4) // [1 2 3 7 4]
//	c := append(a, 5) // [1 2 3 7 5]
//
//	fmt.Println(b) // [1 2 3 7 4]
//	fmt.Println(c) // [1 2 3 7 5]
//}

//=================================================================================================
//===========================     Многопоточная   ==================================================
//=================================================================================================

//func main() {
//	ch := make(chan int)
//	wg := &sync.WaitGroup{}
//	wg.Add(3)
//	for i := 0; i < 3; i++ {
//		go func(v int) {
//			defer wg.Done()
//			ch <- v * v
//		}(i)
//	}
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	var sum int
//	for v := range ch {
//		sum += v
//	}
//	fmt.Printf("result: %d\n", sum)
//}

//func main() {
//	a := 5000
//	wg := &sync.WaitGroup{}
//	wg.Add(a) // Указываем, сколько горутин нужно дождаться
//
//	for i := 0; i < a; i++ {
//		go func(num int) {
//			defer wg.Done() // Уменьшаем счетчик при завершении
//			fmt.Println(num)
//		}(i) // Важно передавать `i` как аргумент, иначе будет race condition!
//	}
//
//	wg.Wait() // Ждём, пока все горутины завершатся
//}

//Будет ошибка что все горутины заблокированы. Какие горутины будут заблокированы? И почему?

//func main() {
//	ch := make(chan int)
//	ch <- 1
//	go func() {
//		fmt.Println(<-ch)
//	}()
//}

//func main() {
//	ch := make(chan bool)
//	ch2 := make(chan bool)
//	ch3 := make(chan bool)
//	go func() {
//		ch <- true
//	}()
//	go func() {
//		ch2 <- true
//	}()
//	go func() {
//		ch3 <- true
//	}()
//
//	select {
//	case <-ch:
//		fmt.Printf("val from ch")
//	case <-ch2:
//		fmt.Printf("val from ch2")
//	case <-ch3:
//		fmt.Printf("val from ch3")
//	}
//}

//var globalMap = map[string][]int{"test1": make([]int, 0), "test2": make([]int, 0), "test3": make([]int, 0)}
//
////var a = 0
//
//func main() {
//	wg := sync.WaitGroup{}
//	wg.Add(3)
//	var a = 0 // добавил
//	go func(a int) {
//		wg.Done()
//		a = 10
//		globalMap["test1"] = append(globalMap["test1"], a)
//	}(a)
//	go func(a int) {
//		wg.Done()
//		a = 11
//		globalMap["test2"] = append(globalMap["test2"], a)
//	}(a)
//	go func(a int) {
//		wg.Done()
//		a = 12
//		globalMap["test3"] = append(globalMap["test3"], a)
//	}(a)
//	wg.Wait()
//	fmt.Printf("%v", globalMap)
//	fmt.Printf("%d", a)
//}

//// Нужно реализовать функцию, которая выполняет поиск query во всех переданных SearchFunc
//// Когда получаем первый успешный результат - отдаем его сразу. Если все SearchFunc отработали
//// с ошибкой - отдаем последнюю полученную ошибку
//
//type Result struct{}
//
//type SearchFunc func(ctx context.Context, query string) (Result, error)
//
//func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
//	var lstError error
//
//	var res Result
//
//	for _, val := range sfs {
//		select {
//		case <-ctx.Done():
//			return Result{}, nil
//		default:
//			res, lstError = val(ctx, query)
//			if lstError != nil {
//				return Result{}, lstError
//			}
//		}
//	}
//	return res, lstError
//}

//// что выведет и как исправить?(уже исправил)
//func main() {
//	m := make(chan string, 3)
//	cnt := 5
//	wg := sync.WaitGroup{}
//	for i := 0; i < cnt; i++ {
//		wg.Add(1)
//		go func(id int) {
//			defer wg.Done()
//			m <- fmt.Sprintf("Goroutine %d", id)
//		}(i)
//	}
//	for i := 0; i < cnt; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			ReceiveFromCh(m)
//		}()
//	}
//
//	wg.Wait()
//}
//func ReceiveFromCh(ch chan string) {
//	fmt.Println(<-ch)
//}

//1. Merge n channels
//2. Если один из входных каналов закрывается, то нужно закрыть все остальные каналы

//func case3(channels ...chan int) chan int {
//	res := make(chan int)
//	wg := sync.WaitGroup{}
//	ctx, cancel := context.WithCancel(context.Background())
//
//	for _, ch := range channels {
//		wg.Add(1)
//		go func(c chan int) {
//			defer wg.Done()
//			for {
//				select {
//				case <-ctx.Done():
//					return
//				default:
//					if v, ok := <-c; !ok {
//						cancel()
//						return
//					} else {
//						select {
//						case <-ctx.Done():
//							return
//						case res <- v:
//						}
//					}
//				}
//			}
//		}(ch)
//	}
//
//	go func() {
//		wg.Wait()
//		close(res)
//	}()
//
//	return res
//}

//Задача 12
//+++++++++++
//1. Конурентно по батчам запросить данные и заь не более chunkSize запросов.писать в файл. Нужна общая конструкция, функции которые делают запрос к сайту и выгрузку в файл можно не реализовывать.
//2. Сделать так, чтобы одновременно выполнялос

//package main
//
//import (
//	"context"
//	"fmt"
//	"net/http"
//	"sync"
//
//	"golang.org/x/sync/errgroup"
//)
//
//const (
//	url       = "http://jsonplaceholder.typicode.com/todos/%d"
//	chunkSize = 100
//	dataCount = 2 << 10
//)
//
//type Resp struct {
//	TaskID  int
//	ErrCode int
//}
//
//func main() {
//	g, ctx := errgroup.WithContext(context.Background())
//	var file sync.Map
//
//	tasks := make(chan int, chunkSize)
//	results := make(chan Resp, chunkSize)
//
//	g.Go(func() error {
//		defer close(tasks)
//		for i := 1; i <= dataCount; i++ {
//			select {
//			case <-ctx.Done():
//				return ctx.Err()
//			case tasks <- i:
//			}
//		}
//		return nil
//	})
//
//	for i := 0; i < chunkSize; i++ {
//		g.Go(func() error {
//			for {
//				select {
//				case <-ctx.Done():
//					return nil
//				case taskID, ok := <-tasks:
//					if !ok {
//						return nil
//					}
//					resp, err := fetch(ctx, taskID)
//					if err != nil {
//						return fmt.Errorf("task %d failed: %w", taskID, err)
//					}
//					select {
//					case <-ctx.Done():
//						return nil
//					case results <- resp:
//					}
//				}
//			}
//		})
//	}
//
//	g.Go(func() error {
//		for {
//			select {
//			case <-ctx.Done():
//				return nil
//			case res, ok := <-results:
//				if !ok {
//					return nil
//				}
//				file.Store(res.TaskID, res.ErrCode)
//				if res.ErrCode != 0 {
//					return fmt.Errorf("error detected for task %d: %d", res.TaskID, res.ErrCode)
//				}
//			}
//		}
//	})
//
//	if err := g.Wait(); err != nil {
//		fmt.Printf("Processing stopped: %v\n", err)
//		return
//	}
//
//	file.Range(func(key, value interface{}) bool {
//		fmt.Printf("Taxi ID: %d, Status: %d\n", key, value)
//		return true
//	})
//}
//
//func fetch(ctx context.Context, id int) (Resp, error) {
//	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(url, id), nil)
//	if err != nil {
//		return Resp{TaskID: id, ErrCode: 500}, err
//	}
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return Resp{TaskID: id, ErrCode: 503}, err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode >= 400 {
//		return Resp{TaskID: id, ErrCode: resp.StatusCode},
//			fmt.Errorf("HTTP error: %d", resp.StatusCode)
//	}
//
//	return Resp{TaskID: id, ErrCode: 0}, nil
//}

//Задача 13
//+++++++++++
//1. Запросить параллельно данные из источников. Если все где-то произошла ошибка, то вернуть ошибку, иначе вернуть nil.
//2. Представим, что теперь функция должна возвращать результат int. Есть функция resp.Size(), для каждого url надо проссумировать и вернуть, если ошибок не было. Просто описать подход к решению
//3. Что делать, если урлов у нас миллионы?

//package main
//
//import (
//	"context"
//	"fmt"
//	"golang.org/x/sync/errgroup"
//	"net/http"
//	"sync"
//)
//
//func main() {
//
//	urls := []string{
//		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
//		"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
//		"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
//		"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
//		"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
//	}
//	size := make(map[string]int64, len(urls))
//
//	err := download(urls, size)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(size)
//}
//
//func download(urls []string, size map[string]int64) error {
//	g, ctx := errgroup.WithContext(context.Background())
//
//	var mu sync.Mutex
//
//	sem := make(chan struct{}, 100)
//
//	for _, url := range urls {
//		url := url
//		sem <- struct{}{}
//		g.Go(func() error {
//			defer func() { <-sem }()
//			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
//			if err != nil {
//				return err
//			}
//
//			resp, err := http.DefaultClient.Do(req)
//			if err != nil {
//				return err
//			}
//			defer resp.Body.Close()
//
//			if resp.StatusCode != http.StatusOK {
//				return fmt.Errorf("HTTP error: %s", resp.Status)
//			}
//
//			mu.Lock()
//			size[url] = resp.ContentLength
//			mu.Unlock()
//
//			return nil
//
//		})
//	}
//
//	if err := g.Wait(); err != nil {
//		return err
//	}
//
//	return nil
//}

//+++++++++++
//Задача 14
//+++++++++++
//1. Что выведет на экран и сколько времени будет работать? миллисекунду?
//2. Нужно ускорить, чтобы работало быстрее. Сколько будет работать теперь? атомики?
//3. Если бы в networkRequest выполнялся реальный сетевой вызов, то какие с какими проблемами мы могли бы столкнуться в данном коде?
//4. Если url немного, а запросов к ним много, то как можно оптимизировать?

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"sync/atomic"
//	"time"
//)
//
//const numRequests = 10000
//
//var count atomic.Int32
//
//var m sync.Mutex
//
//func networkRequest(sem chan struct{}) {
//	defer func() { <-sem }()
//	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
//	count.Add(1)
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	sem := make(chan struct{}, 100)
//
//	wg.Add(numRequests)
//	for i := 0; i < numRequests; i++ {
//		sem <- struct{}{}
//		go func() {
//			defer wg.Done()
//			networkRequest(sem)
//		}()
//	}
//
//	wg.Wait()
//	fmt.Println(count.Load())
//}

//
//Задача 15
//+++++++++++
//
//// Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
//// Её тело нельзя изменять (представим, что внутри сетевой запрос).
//
//// Нужно написать обертку predictableFunc,
//// которая будет работать с заданным фиксированным таймаутом (например, 1 секунду).

//package main
//
//import (
//	"math/rand"
//	"time"
//)
//
//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
//
//// Есть функция, работающая неопределенно долго и возвращающая число.
//// Её тело нельзя изменять (представим, что внутри сетевой запрос).
//func unpredictableFunc() int64 {
//	rnd := rand.Int63n(5000)
//	time.Sleep(time.Duration(rnd) * time.Millisecond)
//
//	return rnd
//}

//// Нужно изменить функцию обертку, которая будет работать с заданным таймаутом (например, 1 секунду).
//// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
//// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
////
//// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
//// Сигнатуру функцию обёртки менять можно.
//func predictableFunc() (int64, error) {
//
//	results := make(chan int64, 1)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
//	defer cancel()
//
//	start := time.Now()
//
//	go func() {
//		results <- unpredictableFunc()
//	}()
//
//	select {
//	case <-ctx.Done():
//		return 0, fmt.Errorf("timout cancel")
//	case result := <-results:
//		fmt.Println(time.Since(start))
//		return result, nil
//	}
//}
//
//func main() {
//	fmt.Println("started")
//
//	fmt.Println(predictableFunc())
//}

////+++++++++++
////Задача 16
////+++++++++++
////
////// Написать код функции, которая делает merge N каналов. Весь входной поток перенаправляется в один канал.
//
//package main
//
//import (
//	"context"
//	"sync"
//)
//
//func merge(cs ...<-chan int) <-chan int {
//	res := make(chan int)
//
//	wg := sync.WaitGroup{}
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	for _, ch := range cs {
//		wg.Add(1)
//		go func(c <-chan int) {
//			defer wg.Done()
//			for {
//				select {
//				case <-ctx.Done():
//					return
//				case val, ok := <-c:
//					if !ok {
//						cancel()
//						return
//					}
//					select {
//					case <-ctx.Done():
//						return
//					case res <- val:
//					}
//				}
//			}
//		}(ch)
//	}
//
//	go func() {
//		wg.Wait()
//		close(res)
//	}()
//
//	return res
//}

////++++++++++
////Задача 17
////+++++++++++
////1. Что выведется? Исправь проблему
////
////# Вариант1
//
//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func main() {
//	x := make(map[int]int, 1)
//
//	var mu sync.Mutex
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 2
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 7
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 10
//	}()
//	time.Sleep(100 * time.Millisecond)
//	fmt.Println("x[1] =", x[1])
//}

////+++++++++++
////Задача 18
////+++++++++++
////1. Иногда приходят нули. В чем проблема? Исправь ее
////2. Если функция bank_network_call выполняется 5 секунд, то за сколько выполнится balance()? Как исправить проблему?
////3. Представим, что bank_network_call возвращает ошибку дополнительно. Если хотя бы один вызов завершился с ошибкой, то balance должен вернуть ошибку.
//
//package main
//
//import "sync"
//
//func balance() (int, error) {
//	x := make(map[int]int, 5)
//	var m sync.Mutex
//	var wg sync.WaitGroup
//	errCh := make(chan error, 5) // Канал для ошибок
//	wg.Add(5)
//
//	for i := 0; i < 5; i++ {
//		i := i
//		go func() {
//			defer wg.Done()
//
//			b, err := bank_network_call(i)
//			if err != nil {
//				errCh <- err
//				return
//			}
//
//			m.Lock()
//			x[i] = b
//			m.Unlock()
//		}()
//	}
//
//	// Закрываем канал ошибок после завершения всех горутин
//	go func() {
//		wg.Wait()
//		close(errCh)
//	}()
//
//	// Проверяем ошибки
//	for err := range errCh {
//		if err != nil {
//			return 0, err
//		}
//	}
//
//	// Считаем сумму
//	sum := 0
//	for _, v := range x {
//		sum += v
//	}
//	return sum, nil
//}

//================================					==================
//================================    Интерфейсы    ==================
//================================					==================

////++++++++++
////Задача 1
////++++++++++
////Что выведет код?
//
//package main
//
//import "fmt"
//
//type impl struct{}
//
//type I interface {
//	C()
//}
//
//func (*impl) C() {}
//
//func A() I {
//	return nil
//}
//
//func B() I {
//	var ret *impl
//	return ret
//}
//
//func main() {
//	a := A()
//	b := B()
//	fmt.Println(a == b)
//}

////++++++++++
////Задача 2
////++++++++++
////1. Добавить код, который выведет тип переменной whoami
//
//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//func printType(whoami interface{}) {
//	fmt.Println(reflect.TypeOf(whoami))
//	fmt.Printf("%T\n", whoami)
//}
//
//func main() {
//	printType(42)
//	printType("im string")
//	printType(true)
//}

//+++++++++
//Задача 3
//++++++++++
//Исправить функцию, чтобы она работала. Сигнатуру менять нельзя

//package main
//
//import "fmt"
//
//func printNumber(ptrToNumber interface{}) {
//	if ptrToNumber == nil {
//		fmt.Println("nil")
//	}
//
//	num, ok := ptrToNumber.(*int)
//	if !ok {
//		fmt.Println("is not a number")
//	} else {
//		if num != nil {
//			fmt.Println(*num)
//		} else {
//			fmt.Println(nil)
//		}
//	}
//}
//
//func main() {
//	v := 10
//	printNumber(&v)
//	var pv *int
//	printNumber(pv)
//	pv = &v
//	printNumber(pv)
//}

//============================					==================
//============================    Разные темы   ==================
//============================					==================

//// промежутки, объединение интервалов
//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func MergeIntervals(intervals [][]int) [][]int {
//	if len(intervals) <= 1 {
//		return intervals
//	}
//
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][0] < intervals[j][0]
//	})
//
//	merged := [][]int{intervals[0]}
//
//	for i := 1; i < len(intervals); i++ {
//		last := merged[len(merged)-1]
//		current := intervals[i]
//
//		if current[0] <= last[1] {
//			if current[1] > last[1] {
//				last[1] = current[1]
//			}
//		} else {
//			merged = append(merged, current)
//		}
//	}
//
//	return merged
//}
//
//func main() {
//	intervals := [][]int{
//		{9, 13},
//		{12, 13},
//		{15, 17},
//		{16, 18},
//		{20, 22},
//		{21, 21},
//	}
//
//	merged := MergeIntervals(intervals)
//
//	fmt.Println("Исходные интервалы:")
//	for _, iv := range intervals {
//		fmt.Printf("[%d, %d] ", iv[0], iv[1])
//	}
//
//	fmt.Println("\nОбъединенные интервалы:")
//	for _, iv := range merged {
//		fmt.Printf("[%d, %d] ", iv[0], iv[1])
//	}
//}

