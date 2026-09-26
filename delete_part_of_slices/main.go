/*
функция DeletingFromSlice, которая принимает слайс целых чисел и возвращает новый слайс целых чисел.
Функция создает независимую копию принятого слайса, без лишней вместимости. Все последующие действия и проверки выполняются исключительно с этой копией.

Функция должна последовательно выполнить со слайсом-копией следующие действия:

Удалить последнее значение, если оно существует и больше 10.
Удалить значение по индексу 2, если такой индекс есть в копии и вместимость копии больше 5.
Удалить первое значение, если оно присутствует и были выполнены оба удаления из пунктов 1 и 2.
Убрать лишнюю вместимость у слайса.
Функция должна вернуть полученный слайс.
*/



package main

import (
	"fmt"
)

func main() {

	sliceX := []int{1, 2, 3, 4, 5, 11, 12}
	fmt.Println(DeletingFromSlice(sliceX))

	//[1000, 1, 2, 6, 11, 100, 8, 500]
}

func DeletingFromSlice(slice1 []int) []int {

	resultSlice := make([]int, len(slice1))
	copy(resultSlice, slice1)

	//Удалить последнее значение, если оно существует и больше 10.
	//Удалить значение по индексу 2, если такой индекс есть в копии и вместимость копии больше 5.
	//Удалить первое значение, если оно присутствует и были выполнены оба удаления из пунктов 1 и 2.
	//Убрать лишнюю вместимость у слайса.
	startLength := len(resultSlice)
	if len(resultSlice) != 0 {
		if resultSlice[len(resultSlice)-1] > 10 {
			resultSlice = resultSlice[:len(resultSlice)-1]
		}
		if len(resultSlice) >= 3 && cap(resultSlice) > 5 {
			resultSlice = append(resultSlice[:2], resultSlice[3:]...)
		}
		if len(resultSlice)+2 == startLength {
			resultSlice = resultSlice[1:]
		}
		resultSlice = resultSlice[:len(resultSlice):len(resultSlice)]
	}

	return resultSlice

}
