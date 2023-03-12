package main

import "fmt"

func main() {
	pricelist := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	total := 0

	fmt.Println("Products list: ")

	for k, v := range pricelist {
		if v > 500 {
			fmt.Println(k)
		}
	}

	for _, v := range order {
		total += pricelist[v]
	}
	fmt.Println("Стоимость заказа ", total)

	find([]int{5, 3, 2, 12, 6, 8, 23, 54, 33}, 20)

}

func find(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
}

// как можно заметить, алгоритм пройдётся по массиву всего один раз
// если бы мы искали подходящее значение каждый раз через перебор массива, то пришлось бы сделать гораздо больше вычислений
