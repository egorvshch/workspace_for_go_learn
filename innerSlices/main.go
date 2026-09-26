/*
Функция mirrorMatrix, которая принимает двумерный слайс целых чисел и возвращает новый двумерный слайс.
В возвращаемом слайсе каждая строка должна быть отзеркалена по горизонтали, то есть элементы внутри каждой строки должны располагаться в обратном порядке.
*/

package main

import (
	"fmt"
)

func main() {

	slice2 := [][]int{
		{10, 9, 8, 7},
		{10, 9, 8, 7},
		{10, 9, 8, 7},
	}
	fmt.Println(mirrorMatrix(slice2))
}
func mirrorMatrix(matrix [][]int) [][]int { 
	// Ваш код
	mirrorX := make([][]int, len(matrix))
	for i := range mirrorX {
		mirrorX[i] = []int{}
	}
	for i, innerSlice := range matrix {
		for k := range innerSlice {
			mirrorX[i] = append(mirrorX[i], matrix[i][k])
		}

	}

	for i, innerSlice := range mirrorX {
		for k := 0; k <= len(innerSlice)/2-1; k++ {
			mirrorX[i][k], mirrorX[i][len(innerSlice)-1-k] = mirrorX[i][len(innerSlice)-1-k], mirrorX[i][k]
		}

	}
	return mirrorX
