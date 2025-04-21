//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//// in [0, 0, 3] [1, 2, 5, 7]
//// out [0, 0, 1, 2, 3, 5, 7]
//func foo(ch1, ch2 chan int, wg *sync.WaitGroup) chan int {
//	out := make(chan int)
//
//	val1, ok1 := <-ch1
//	val2, ok2 := <-ch2
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for ok1 || ok2 {
//			if !ok2 || (ok1 && val1 <= val2) {
//				out <- val1
//				val1, ok1 = <-ch1
//			} else {
//				out <- val2
//				val2, ok2 = <-ch2
//			}
//		}
//	}()
//
//	go func() {
//		wg.Wait()
//		close(out)
//	}()
//
//	return out
//}
//
//func main() {
//	ch := make(chan int, 3)
//	ch <- 0
//	ch <- 0
//	ch <- 3
//	close(ch)
//	ch1 := make(chan int, 4)
//	ch1 <- 1
//	ch1 <- 2
//	ch1 <- 5
//	ch1 <- 7
//	close(ch1)
//	wg := sync.WaitGroup{}
//
//	res := foo(ch1, ch, &wg)
//
//	for val := range res {
//		fmt.Println(val)
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	s1, s2 := []int{1, 2, 3}, []int{4, 5, 6, 7, 8}
//	fmt.Println(zip(s1, s2)) // [[1 4] [2 5] [3 6]]
//}
//func zip(s1 []int, s2 []int) [][]int {
//	if len(s1) == 0 && len(s2) == 0 {
//		return [][]int{}
//	}
//
//	minLen := len(s1)
//	maxLen := len(s1)
//	start := s1[0]
//
//	if minLen > len(s2) {
//		minLen = len(s2)
//		maxLen = len(s1)
//		start = s2[0]
//	}
//
//	res := make([][]int, 0, maxLen)
//
//	for i := start - 1; i < maxLen; i++ {
//		res = append(res, []int{s1[i], s2[i]})
//	}
//
//	return res
//
//}

//package main
//
//import "fmt"
//
//func foo(sl []int) []int {
//	if len(sl) == 0 {
//		return nil
//	}
//
//	res := make([]int, 0, len(sl))
//
//	m := make(map[int]bool)
//
//	minVal, maxVal := sl[0], sl[0]
//
//	for _, val := range sl {
//		if val < minVal {
//			minVal = val
//		}
//		if val > maxVal {
//			maxVal = val
//		}
//	}
//
//	for _, val := range sl {
//		m[val] = true
//	}
//
//	for i := minVal; i < maxVal; i++ {
//		if !m[i] {
//			res = append(res, i)
//		}
//	}
//
//	return res
//}
//
//func main() {
//	arr := []int{2, 1, 5, 4, 8, 7}
//	fmt.Println(foo(arr))
//}

//package main
//
//import (
//	"fmt"
//)
//
//func dub(sl1, sl2 []int) []int {
//	lena := len(sl1)
//	if lena < len(sl2) {
//		lena = len(sl2)
//	}
//
//	res := make([]int, 0, lena)
//
//	counter := make(map[int]bool)
//
//	for _, val := range sl1 {
//		counter[val] = true
//	}
//
//	for _, val := range sl2 {
//		if counter[val] {
//			res = append(res, val)
//		}
//	}
//
//	return res
//
//}
//
//func main() {
//	arr := []int{1, 2, 3, 4}
//	arr1 := []int{56, 768, 3, 6, 4}
//	fmt.Println(dub(arr, arr1))
//}
//
//package main
//
//func main() {
//	f()
//}
//
//func f() {
//	var m = make([]int, 500)
//	wg := new(sync.WaitGroup)
//	for i := 0; i < 500; i++ {
//		wg.Add(1)
//		go func(i int) {
//			m[i] = i
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	fmt.Print("f done ")
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := []int{0, 0}
//
//	if a == nil {
//		fmt.Println("true")
//	} else {
//		fmt.Println("false")
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	b := new(bool)
//	modify(b)
//	fmt.Println(b)
//}
//
//func modify(b *bool) {
//	b = nil
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//	counter := 0
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			defer mu.Unlock()
//			mu.Lock()
//			counter++
//		}()
//	}
//	wg.Wait()
//	fmt.Println(counter)
//}

//package main
//
//import "fmt"
//
//func reverseSliceInPlace(sl []int) []int { 1 2 3 4 5 6
//	for i := 0; i < len(sl)/2; i++ {
//		j := len(sl) - 1 - i
//		sl[i], sl[j] = sl[j], sl[i]
//	}
//	return sl
//}
//
//func main() {
//	arr := []int{1, 2, 3, 4, 5}
//	fmt.Println(reverseSliceInPlace(arr))
//}

//package main
//
//import (
//	"fmt"
//)
//
//type Foo struct{}
//
//type Reader interface {
//	MyRead()
//}
//
//type Writer interface {
//	MyWrite()
//}
//
//func (f *Foo) MyRead() {}
//
//func (f *Foo) MyWrite() {
//	fmt.Println("Way")
//}
//
//func main() {
//
//	var myreader Reader = &Foo{}
//
//	mywriter := myreader.(Writer)
//
//	mywriter.MyWrite()
//
//}

// -------Секция Озон------------------------------------------------

// что выведет программа?
//package main
//
//import "fmt"
//
//func main() {
//	nums := []int{1, 2, 3}
//	fmt.Println(len(nums), cap(nums)) // len 3 cap 3
//
//	addNum(nums[0:2])
//	fmt.Println(nums)                 // 1 2 4
//	fmt.Println(len(nums), cap(nums)) // len 3 cap 3
//
//	addNums(nums[0:2])
//	fmt.Println(nums)                 //
//	fmt.Println(len(nums), cap(nums)) // len 3 cap 3
//
//}
//
//func addNum(nums []int) {
//	nums = append(nums, 4)
//}
//
//func addNums(nums []int) {
//	nums = append(nums, 5, 6)
//	fmt.Println(nums)
//}

// Функция для попарного объединения двух слайсов
//package main
//
//import "fmt"
//
//// Функция для попарного объединения двух слайсов попарно
//func mergeSlicesPairwise(slice1, slice2 []int) [][]int {
//	// Определяем длину минимального слайса
//	minLen := len(slice1)
//	if len(slice2) < minLen {
//		minLen = len(slice2)
//	}
//
//	// Создаем результирующий слайс
//	result := make([][]int, minLen)
//
//	// Заполняем результирующий слайс парами
//	for i := 0; i < minLen; i++ {
//		result[i] = []int{slice1[i], slice2[i]}
//	}
//
//	return result
//}
//
//func main() {
//	// Пример использования
//	slice1 := []int{1, 2, 3, 4}
//	slice2 := []int{10, 20, 30}
//
//	merged := mergeSlicesPairwise(slice1, slice2)
//	fmt.Println(merged) // Вывод: [[1 10] [2 20] [3 30]]
//}

//// Функция для генерации слайса из n уникальных случайных чисел
//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func generateUniqueRandomNumbers(n, max int) ([]int, error) {
//	// Проверка на корректность входных данных
//	if n > max {
//		return nil, fmt.Errorf("n не может быть больше max")
//	}
//
//	// Инициализация генератора случайных чисел
//	rand.Seed(time.Now().UnixNano())
//
//	// Используем map для отслеживания уникальных чисел
//	uniqueNumbers := make(map[int]bool)
//	result := make([]int, 0, n)
//
//	// Генерация n уникальных чисел
//	for len(result) < n {
//		num := rand.Intn(max) // Генерация случайного числа от 0 до max-1
//		if !uniqueNumbers[num] {
//			uniqueNumbers[num] = true
//			result = append(result, num)
//		}
//	}
//
//	return result, nil
//}
//
//func main() {
//	// Пример использования
//	n := 10
//	max := 100
//
//	numbers, err := generateUniqueRandomNumbers(n, max)
//	if err != nil {
//		fmt.Println("Ошибка:", err)
//		return
//	}
//
//	fmt.Println("Уникальные случайные числа:", numbers)
//}

//// что выведет
//package main
//
//import "fmt"
//
//func main() {
//	// Создаем срез с начальной емкостью 5 для избежания реаллоцирования
//	x := []int{} // len 0 cap 0
//	fmt.Println(len(x), cap(x))
//	x = append(x, 0) // len 1 cap 1
//	fmt.Println(len(x), cap(x))
//	x = append(x, 1) // len 2 cap 2
//	fmt.Println(len(x), cap(x))
//	x = append(x, 2) // len 3 cap 4
//	fmt.Println(len(x), cap(x))
//	// 0 1 2
//	// Добавляем элементы в x — изменения отразятся на всех срезах
//	y := append(x, 3) // len 4 cap 4
//	z := append(x, 4) // len 4 cap 4
//
//	// Все три среза указывают на общий массив и видят изменения
//	fmt.Println(x, y, z) // Вывод: [0 1 2] [0 1 2 4] [0 1 2 4]
//}

//// написать функцию которая выводит дубликаты
//package main
//
//import "fmt"
//
//func foo(sl1 []int, sl2 []int) []int {
//	res := make([]int, 0)
//
//	uMap := make(map[int]struct{})
//
//	for _, val := range sl1 {
//		uMap[val] = struct{}{}
//	}
//
//	for _, val := range sl2 {
//		if _, ok := uMap[val]; !ok {
//			res = append(res, val)
//		}
//	}
//
//	return res
//}
//
//func main() {
//	sl1 := []int{1, 2, 3, 6}
//	sl2 := []int{3, 4, 6}
//	fmt.Println(foo(sl1, sl2))
//}

//// нужно найти число которое не хватает
//
//package main
//
//import (
//	"fmt"
//)
//
//func foo(nums []int) int {
//	uMap := make(map[int]struct{})
//
//	for _, val := range nums {
//		uMap[val] = struct{}{}
//	}
//
//	min := nums[0]
//	max := nums[0]
//
//	for _, val := range nums {
//		if val < min {
//			min = val
//		}
//		if val > max {
//			max = val
//		}
//	}
//
//	var res int
//
//	for i := min; i <= max; i++ {
//		if _, ok := uMap[i]; !ok {
//			res = i
//		}
//	}
//	return res
//}
//
//func main() {
//	nums := []int{1, 2, 3, 4, 6, 5, 8, 7, 10}
//	fmt.Println(foo(nums))
//}

//Дана таблица "orders":
//
//+-----+-----------+------------------------+--------------+
//|  id |  user_id  |             created_at |  price_total |
//+-----+-----------+------------------------+--------------+
//|   1 |       111 |    2024-01-01 10:00:00 |        2 000 |
//|   2 |       222 |    2024-01-02 10:00:00 |          100 |
//|   3 |       111 |    2024-04-01 10:00:00 |       20 000 |
//|   4 |       222 |    2024-05-01 10:00:00 |        5 000 |
//|   5 |       333 |    2024-05-02 10:00:00 |       10 000 |
//+-----+-----------+------------------------+--------------+
//
//
//Напишите SQL-запрос, который вернёт количество заказов по каждому пользователю с price_total больше или равным 1000 в таком виде отсортированному по количеству заказов в обратном порядке:
//
//+----------+-------------+
//|  user_id | order_count |
//+----------+-------------+
//|      111 |           2 |
//|      222 |           1 |
//|      333 |           1 |
//+----------+-------------+

//SELECT
//user_id,
//COUNT(*) AS order_count
//FROM
//orders
//WHERE
//price_total >= 1000
//GROUP BY
//user_id
//ORDER BY
//order_count DESC;

// написать функцию которая вычисляет длину самого большого паллиндрома который возможен из строки
//package main
//
//import (
//	"fmt"
//)
//
//func longestPalindrome(s string) int {
//	count := make(map[rune]int)
//	for _, char := range s {
//		count[char]++
//	}
//
//	length := 0
//	oddFound := false
//	for _, freq := range count {
//		if freq%2 == 0 {
//			length += freq
//		} else {
//			length += freq - 1
//			oddFound = true
//		}
//	}
//
//	if oddFound {
//		length++
//	}
//
//	return length
//}
//
//func main() {
//	input := "aaabbbccccdd"
//	result := longestPalindrome(input)
//	fmt.Printf("Output: %d (палиндром может быть построен из этих символов)\n", result)
//}

// надо написать sql запрос который выводит имена файлов, обработка которых еще не закончена

//SELECT d.name
//FROM documents d
//JOIN statuses s ON d.id = s.document_id
//JOIN (
//SELECT
//document_id,
//MAX(created_at) AS last_created_at
//FROM statuses
//GROUP BY document_id
//) last_status ON s.document_id = last_status.document_id AND s.created_at = last_status.last_created_at
//WHERE s.status != 'completed';

// написать функцию которая выводит ошибку не использую сторонних пакетов

//package main
//
//// Кастомный тип ошибки
//type MyError struct {
//	msg string
//}
//
//// Метод для реализации интерфейса error
//func (e *MyError) Error() string {
//	return e.msg
//}
//
//// Функция, которая возвращает ошибку
//func (e *MyError) handle() error {
//	return e // Возвращаем саму структуру, так как она реализует интерфейс error
//}
//
//func main() {
//	// Создаем экземпляр ошибки
//	s := &MyError{"error"}
//
//	// Вызываем метод handle и проверяем ошибку
//	err := s.handle()
//	if err != nil {
//		println("Ошибка:", err.Error()) // Выводим ошибку
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	s := "test"
//	s1 := []rune(s)
//	fmt.Println(s1[0]) // Выведет: 116
//
//	s1[0] = 'R'     // Исправлено: используем одинарные кавычки для символа
//	fmt.Println(s1) // Выведет: [82 101 115 116]
//}

// ----------------------------------- Тех Собес Озон -----------------------------------------------------------------------

//// убрать все нули из массива
//package main
//
//import "fmt"
//
//func main() {
//	sl := []int{0, 1, 2, 3, 0, 4, 5, 0}
//	deletenull(&sl)
//	fmt.Println(sl)
//}
//
//func deletenull(sl *[]int) {
//	j := 0
//	for i := 0; i < len(*sl); i++ {
//		if (*sl)[i] != 0 {
//			(*sl)[j] = (*sl)[i]
//			j++
//		}
//	}
//	*sl = (*sl)[:j]
//}

//// у нас есть две строки: надо сравнить одинаковы ли эти строки, # - означает backspace
//package main
//
//import "fmt"
//
//func isEqual(s1, s2 string) bool {
//	run1 := []rune{}
//	run2 := []rune{}
//
//	for _, char := range s1 {
//		if char == '#' {
//			run1 = run1[:len(run1)-1]
//		} else {
//			run1 = append(run1, char)
//		}
//	}
//
//	for _, char := range s2 {
//		if char == '#' {
//			run2 = run2[:len(run2)-1]
//		} else {
//			run2 = append(run2, char)
//		}
//	}
//
//	if string(run1) == string(run2) {
//		return true
//	}
//
//	return false
//
//}
//
//func main() {
//	s1 := "ab#c"
//	s2 := "ad#c"
//	fmt.Println(isEqual(s1, s2))
//}

//// даны два массива, один отсортированный и дальше нули, другой тоже отсортированный, надо их смержить
//package main
//
//import "fmt"
//
//func merge(sl1 []int, m int, sl2 []int, n int) []int {
//	res := make([]int, 0)
//
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	go func() {
//		for i := 0; i < m; i++ {
//			if sl1[i] != 0 {
//				ch1 <- sl1[i]
//			}
//		}
//		close(ch1)
//	}()
//
//	go func() {
//		for i := 0; i < n; i++ {
//			ch2 <- sl2[i]
//		}
//		close(ch2)
//	}()
//
//	val1, ok1 := <-ch1
//	val2, ok2 := <-ch2
//	for ok1 || ok2 {
//		if !ok2 || (ok1 && val1 <= val2) {
//			res = append(res, val1)
//			val1, ok1 = <-ch1
//		} else {
//			res = append(res, val2)
//			val2, ok2 = <-ch2
//		}
//	}
//
//	return res
//}
//
//func main() {
//	arr1 := []int{1, 3, 5, 7, 0, 0, 0, 0}
//	arr2 := []int{2, 4, 6, 8}
//	fmt.Println(merge(arr1, 4, arr2, 4))
//}

//// составить бд для библиотеки, есть 3 сущности: Автор, Книга и Читатель.
//// у одного читателя может быть только одна книга.
//
//CREATE TABLE Author (
//	id serial
//	name varchar not null
//	)
//
//CREATE TABLE Book (
//id serial
//name varchar not null
//	)
//
//CREATE TABLE Reader (
//id serial
//name varchar not null
//	)
//
//CREATE TABLE Book_Order (
//	id serial
//	reader_id int
//	book_id int unique
//	)
//
//CREATE TABLE Book_Author (
//	book_id int
//	author_id int
//)

//SELECT b.name as book_name
//FROM BookOrder bo
//JOIN BOOK b ON b.id = bo.book_id

//SELECT b.name
//FROM Book_Author ba
//JOIN Book b ON b.id = ba.book_id
//GROUP ba.book_id, b.name
//HAVING COUNT(ba.author_id) > 3;


//SELECT a.name
//FROM Book_Order bo
//JOIN Book_Author ba ON ba.book_id = bo.book_id
//JOIN Author a ON a.id = ba.author_id
//GROUP BY ba.author_id, a.name
//ORDER BY bo.reader_id DESC
//LIMIT 3;


//// у нас есть набор url, надо написать программу которая поочередно выполнит запросы, если
//// получили ответ OK, то написать url - ok, иначе utl - not ok
//
//package main
//
//import (
//	"context"
//	"fmt"
//	"net/http"
//	"sync"
//)
//
//type Response struct {
//	StatusCode int
//	Err        error
//	Url        string
//}
//
//func main() {
//	urls := []string{
//		"http://google.com",
//		"qfqwfqwfqwfqwf",
//		"qfqwfqwfqwfqwfqw",
//		"qqwfqwfqwfqwfqwf",
//	}
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	ch := make(chan Response)
//	wg := sync.WaitGroup{}
//	for _, val := range urls {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			select {
//			case <-ctx.Done():
//				return
//			default:
//				resp, err := http.Get(val)
//				ch <- Response{Err: err, StatusCode: resp.StatusCode, Url: val}
//			}
//		}()
//	}
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	for val := range ch {
//		if val.Err != nil || val.StatusCode != 200 {
//			fmt.Printf("%s - ok", val.Url)
//		} else {
//			fmt.Printf("%s - not ok", val.Url)
//		}
//	}
//
//}

//Сервис, который обрабатывает данные в реальном времени и выдает аналитику о количестве запросов
//пользователя за последние 5 минут, можно использовать структуру данных,
//такую как слайс или очередь с фиксированным размером для хранения временных меток запросов пользователя.
//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//type Service struct {
//	mu       sync.Mutex
//	requests map[string][]time.Time
//}
//
//func NewService() *Service {
//	return &Service{
//		requests: make(map[string][]time.Time),
//	}
//}
//
//func (s *Service) HandleEvent(userName string, currentTime time.Time) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	// Добавляем текущее время в список запросов пользователя
//	s.requests[userName] = append(s.requests[userName], currentTime)
//
//	// Очищаем старые запросы, которые были более 5 минут назад
//	s.cleanupOldRequests(userName, currentTime)
//}
//
//func (s *Service) GetCount(userName string, currentTime time.Time) int {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	// Очищаем старые запросы перед подсчетом
//	s.cleanupOldRequests(userName, currentTime)
//
//	// Возвращаем количество запросов за последние 5 минут
//	return len(s.requests[userName])
//}
//
//func (s *Service) cleanupOldRequests(userName string, currentTime time.Time) {
//	threshold := currentTime.Add(-5 * time.Minute)
//	newRequests := []time.Time{}
//
//	for _, t := range s.requests[userName] {
//		if t.After(threshold) {
//			newRequests = append(newRequests, t)
//		}
//	}
//
//	s.requests[userName] = newRequests
//}
//
//func main() {
//	service := NewService()
//
//	// Пример использования
//	userName := "user1"
//	currentTime := time.Now()
//
//	// Обработка нескольких событий
//	for i := 0; i < 10; i++ {
//		service.HandleEvent(userName, currentTime.Add(time.Duration(i)*time.Second))
//	}
//
//	// Получение количества запросов за последние 5 минут
//	count := service.GetCount(userName, currentTime.Add(6*time.Minute))
//	fmt.Printf("Количество запросов за последние 5 минут: %d\n", count)
//}

//// удалить элемент из односвязного списка
//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	dummy := &ListNode{Next: head}
//	first, second := dummy, dummy
//
//	// Сдвигаем первый указатель на n + 1 шагов вперед
//	for i := 0; i <= n; i++ {
//		first = first.Next
//	}
//
//	// Двигаем оба указателя до конца списка
//	for first != nil {
//		first = first.Next
//		second = second.Next
//	}
//
//	// Удаляем n-й элемент с конца
//	second.Next = second.Next.Next
//
//	return dummy.Next
//}

//// задача на релизацию кэша с шардами
//package main
//
//import (
//	"fmt"
//	"hash/fnv"
//	"sync"
//)
//
//type MyCache interface {
//	Get(key string) string
//	Set(key string, value string)
//}
//
//type Cache struct {
//	shards []*Shard
//}
//
//type Shard struct {
//	data map[string]string
//	mu   sync.RWMutex
//}
//
//func NewCache(shardCount int) *Cache {
//	shards := make([]*Shard, shardCount)
//	for idx := range shards {
//		shards[idx] = &Shard{
//			data: make(map[string]string),
//		}
//	}
//	return &Cache{
//		shards: shards,
//	}
//}
//
//func (c *Cache) Set(key string, value string) {
//	shard := c.GetShard(key)
//	shard.mu.Lock()
//	defer shard.mu.Unlock()
//	shard.data[key] = value
//}
//
//func (c *Cache) Get(key string) string {
//	shard := c.GetShard(key)
//	shard.mu.RLock()
//	defer shard.mu.RUnlock()
//	val, ok := shard.data[key]
//	if !ok {
//		return ""
//	}
//	return val
//}
//
//func (c *Cache) GetShard(key string) *Shard {
//	hasher := fnv.New32()
//	_, _ = hasher.Write([]byte(key))
//	hash := hasher.Sum32()1
//
//	return c.shards[hash%uint32(len(c.shards))]
//}
//
//func main() {
//	cache := NewCache(3)
//	cache.Set("one", "two")
//	fmt.Println(cache.Get("one"))
//}

// реализовать LRU CACHE

//package main
//
//import "fmt"
//
//type LRUCache struct {
//	counter []string
//	data    map[string]struct{}
//	limit   int
//}
//
//func NewCache(limit int) *LRUCache {
//	return &LRUCache{
//		counter: make([]string, 0),
//		data:    make(map[string]struct{}),
//		limit:   limit,
//	}
//}
//
//func (l *LRUCache) CountWord(word string) {
//	if _, ok := l.data[word]; !ok {
//		l.counter = append(l.counter, word)
//	}
//	l.data[word] = struct{}{}
//
//	if len(l.counter) > l.limit {
//		delete(l.data, l.counter[0])
//		l.counter = l.counter[1:]
//	}
//
//}
//
//func main() {
//
//	c := NewCache(3)
//
//	words := []string{"bmw", "merc", "yaguar", "rangerover", "buggatti"}
//
//	for _, val := range words {
//		c.CountWord(val)
//	}
//
//	fmt.Println(c.data)
//
//}

//Есть 3 сущности - пользователь, чат, сообщение
//
//- У пользователя есть имя и дата регистрации
//- У чата есть название и дата создания
//- У сообщения есть текст, автор и дата создания
//- Пользователь может состоять в нескольких чатах одновременно
//Сообщение обязательно принадлежит чату, сообщение не может принадлежать более чем 1 чату одновременно
//Нужно описать предметную область в виде таблиц
//
//Доп: Выбрать все чаты пользователя Вася в формате (chat_id,chat_name)

CREATE TABLE users (
user_id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
registration_date TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE chats (
chat_id SERIAL PRIMARY KEY,
chat_name VARCHAR(255) NOT NULL,
creation_date TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE messages (
message_id SERIAL PRIMARY KEY,
text TEXT NOT NULL,
user_id INT REFERENCES users(user_id),
chat_id INT REFERENCES chats(chat_id),
creation_date TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Таблица для связи многие-ко-многим между пользователями и чатами
CREATE TABLE user_chat (
user_id INT REFERENCES users(user_id),
chat_id INT REFERENCES chats(chat_id),
PRIMARY KEY (user_id, chat_id)
);

SELECT c.chat_id, c.chat_name
FROM users u
JOIN user_chat uc ON u.user_id = uc.user_id
JOIN chats c ON uc.chat_id = c.chat_id
WHERE u.name = 'Вася';