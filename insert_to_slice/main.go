package main

import (
	"fmt"
	"slices"
)

func main() {

	var sliceX []int
	fmt.Println(PlayWithSlice(sliceX))

	//[1000, 1, 2, 6, 11, 100, 8, 500]
}

func PlayWithSlice(slice1 []int) []int {

	//Клонирование слайса: Создайте новый слайс, который будет являться клоном переданного слайса. Все дальнейшие операции должны выполняться только с клоном.

	resultSlice := make([]int, len(slice1))
	resultSlice = slices.Clone(slice1)

	//Вставка числа 100: Найдите первое значение с конца клона, которое больше 10. После этого значения вставьте число 100. Если такого значения не найдено, этот шаг можно пропустить.

	value0 := 100

	for i := len(resultSlice) - 1; i >= 0; i-- {

		if resultSlice[i] > 10 {

			before := resultSlice[:i+1]
			after := append([]int{value0}, resultSlice[i+1:]...)
			resultSlice = append(before, after...)

			break
		}
	}

	//Вставка числа 500: Если сумма всех чисел в текущем клоне больше 100, добавьте число 500 в конец слайса.

	summValSlice := 0
	for _, val := range resultSlice {
		summValSlice += val
	}

	if summValSlice > 100 {
		value1 := 500
		resultSlice = append(resultSlice, value1) // добавление число 500 в конец слайса.
	}

	//Вставка числа 1000: Если в оригинальном слайсе четных чисел больше, чем нечетных, вставьте число 1000 в начало клона слайса.

	count1 := 0 // %2 != 0 Не четные
	count2 := 0 // %2 == 0 Четные

	for _, val := range slice1 {
		if val%2 != 0 {
			count1++
		} else {
			count2++
		}
	}

	if count2 > count1 {
		value2 := 1000
		resultSlice = append([]int{value2}, resultSlice...) // вставьте число 1000 в начало клона слайса.
	}

	return resultSlice

}

/*
func insertElement(slice []int, pos int, value int) []int {
    n := len(slice)
    newSlice := make([]int, n+1)
    copy(newSlice[:pos], slice[:pos])          // Копируем левую половину
    newSlice[pos] = value                       // Добавляем новый элемент
    copy(newSlice[pos+1:], slice[pos:n])       // Копируем правую половину
    return newSlice
}

func main() {
    slice := []int{1, 2, 3, 4, 5}
    pos := 2
    value := 100
    result := insertElement(slice, pos, value)
    fmt.Println(result) // Выведет: [1 2 100 3 4 5]
}
*/
