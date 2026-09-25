package main

import (
	"fmt"
)

func main() {

	var massive = [5]int{3, 8, 1, 8, 1}
	fmt.Println(secretGenerator(massive))

}

func secretGenerator(nums [5]int) string {

	var secretKey string
	var newMassive [len(nums) + 2]int
	min := nums[0]
	max := nums[0]
	for i := range nums {

		newMassive[i+1] = nums[i]

		if nums[i] < min {
			newMassive[0] = nums[i]
		}
		if nums[i] > max {
			newMassive[len(newMassive)-1] = nums[i]
		}

	}

	for i := range newMassive {
		switch i {
		case 0, len(newMassive) - 1:
			secretKey += fmt.Sprintf("%d", newMassive[i])
		default:
			if newMassive[i]%2 == 0 {
				secretKey += fmt.Sprintf("E%d", newMassive[i])
			} else {
				secretKey += fmt.Sprintf("O%d", newMassive[i])
			}
		}
	}
	return secretKey
}

// 1O3E8O1E8O18
