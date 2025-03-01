// package main

// import "fmt"

// func main() {
//  a1 := make([]int, 0, 10)
//  a1 = append(a1, []int{1, 2, 3, 4, 5}...)
//  a2 := append(a1, 6)
//  a3 := append(a1, 7)
//  a4 := append(a1, 45, 56, 67, 67)

//  fmt.Println(a1, a2, a3, a4)
// }

package main

import (
	"fmt"
	"math"
	"sort"
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

func FindTarget(goods []int, target int) int {
	low := 0
	high := len(goods) - 1

	for low <= high {
		mid := low + (high-low)/2
		if goods[mid] == target {
			return goods[mid]
		} else if goods[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if high < 0 {
		return goods[low]
	}
	if low >= len(goods) {
		return goods[high]
	}

	if math.Abs(float64(target-goods[low])) < math.Abs(float64(target-goods[high])) {
		return goods[low]
	}
	return goods[high]
}

func FindGood(goods []int, needs []int) int {
	sort.Ints(goods)

	sum := 0
	for _, val := range needs {
		closest := FindTarget(goods, val)
		sum += int(math.Abs(float64(val - closest)))
	}
	return sum
}

func main() {
	goods := []int{8, 3, 5}
	buyerNeeds := []int{5, 14, 12, 44, 55}
	fmt.Println(FindGood(goods, buyerNeeds))
}
