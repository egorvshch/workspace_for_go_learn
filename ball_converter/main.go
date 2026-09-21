// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
)

func main() {
	examBallVal, err := ExamBallConvert()
	if err != nil {
		log.Fatalf("Error caused: %s", err)
	}
	fmt.Println("Ваша буквенная оценка:", examBallVal)
}

func ExamBallConvert() (string, error) {
	var input int
	var ball string
	fmt.Printf("Введите числовую оценку (целое число) в диапазоне от 0 до 100: ")
	_, err := fmt.Scan(&input)
	if err != nil || input > 100 || input < 0 {
		return "", fmt.Errorf("Ошибка ввода числовой оценки\n")
	} else {
		switch {
		case input >= 90 && input <= 100: // 90–100: A
			ball = "A"

		case input >= 80 && input <= 89: // 80–89: B
			ball = "B"

		case input >= 70 && input <= 79: // 70–79: C
			ball = "C"

		case input >= 60 && input <= 69: // 60–69: D
			ball = "D"

		default: // Ниже 60: F
			ball = "F"
		}

		return ball, nil
	}
}
