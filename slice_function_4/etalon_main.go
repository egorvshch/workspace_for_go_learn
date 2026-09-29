package main

import (
	"errors"
	"math"
	"math/rand/v2"
	"slices"
)

func CreateSlice(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n least than zero")
	}

	slice := make([]int, 0, n)

	for range n {
		num := rand.IntN(21) - 10
		slice = append(slice, num)
	}

	return slice, nil

}

func FilterSlice(numbers []int) []int {

	res := make([]int, 0)
	for i := 1; i < len(numbers); i++ {
		if (numbers[i-1] > numbers[i]) && isDivided(numbers[i]) {
			res = append(res, numbers[i])
		}
	}

	return res
}

func isDivided(num int) bool {
	divs := []int{2, 5, 6, 9}
	for _, d := range divs {
		if num%d == 0 {
			return true
		}
	}
	return false
}

func MaxSumWithNegative(numbers []int, k int) []int {

	var sliceMaxSum int
	var resultSlice []int

	for i := 0; i < len(numbers)-k+1; i++ {
		sum := 0
		isNegative := false

		subSlice := numbers[i : i+k]
		for _, num := range subSlice {
			sum += num
			if num < 0 {
				isNegative = true
			}
		}

		if (resultSlice == nil || sum > sliceMaxSum) && isNegative {
			sliceMaxSum = sum
			resultSlice = slices.Clone(subSlice)
		}
	}
	return resultSlice
}

func SortByParity(numbers []int) []int {

	clone := slices.Clone(numbers)

	slices.SortFunc(clone, func(a, b int) int {
		aAbs := math.Abs(float64(a % 2))
		bAbs := math.Abs(float64(b % 2))

		if aAbs != bAbs {
			return int(aAbs - bAbs)
		}

		if a%2 == 0 {
			return b - a
		}
		return a - b
	})
	return clone

}
