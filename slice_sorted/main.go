/*
Функция sortMagic, которая принимает слайс целых чисел и сортирует его в порядке убывания, при этом четные числа должны идти перед нечетными. Ноль будем считать как обычное четное число.
*/

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {

	sliceX := []int{}
	sortMagic(sliceX)
	fmt.Println(sliceX)

	//[1000 500 100 8 6 2 11 1]
}

func sortMagic(slice1 []int) {

	typeNum := func(x int) int {
		if x%2 == 0 {
			return 0
		}
		return 1
	}

	slices.SortFunc(slice1, func(a, b int) int {

		return cmp.Or(cmp.Compare(typeNum(a), typeNum(b)), cmp.Compare(b, a))

	})

}

/*
func sortMagic(slice1 []int) {
	slices.SortFunc(slice1, func(a, b int) int {
		// определяем: число четное или нечетное
		isAEven := a%2 == 0
		isBAEven := b%2 == 0

		// Четные идут раньше нечетных
		if isAEven != isBAEven {
			if !isAEven {
				return 1
			}
			return -1
		}

		// сортируем внутри группы
		return cmp.Compare(b, a)
	})

}
*/
