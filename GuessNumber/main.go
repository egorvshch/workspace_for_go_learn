package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

var attemps int
var random int
var answer string

func main() {
	randomGenegator()
	for {
		fmt.Print("Ваше предположение (либо, для завершения, введите слово \"выход\"): ")

		answer = getInput()

		if answer == "выход" {
			fmt.Printf("Спасибо за игру! До свидания!")
			os.Exit(0)
		}
		answerNum, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Printf("Ошибка ввода\n")
		}
		if err == nil {
			checkNum, checkStr := checkResult(answerNum)

			if checkNum != 0 {
				fmt.Println(checkStr)
			}
			if checkNum == 0 {
				fmt.Println(checkStr)
				fmt.Print("Хотите сыграть еще раз? (если хотите, напишите слово да): ")

				for {
					answer = getInput()
					if answer == "нет" {
						fmt.Printf("Спасибо за игру! До свидания!")
						os.Exit(0)
					}
					if answer == "да" {
						randomGenegator()
						break
					}
					fmt.Println("Ошибка ввода, введите \"да\" или \"нет \"")
					fmt.Print("Хотите сыграть еще раз? (если хотите, напишите слово да): ")
				}

			}
			randomGenegator()

		}
	}

}

// Генерация случайного числа:
func randomGenegator() {
	attemps = 0
	random = rand.IntN(100) + 1
	fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")
	fmt.Println(random)
}

// Функция получения ввода от пользователя:
func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	scannerOut := strings.ToLower(scanner.Text())

	return scannerOut
}

// Проверка результата и отслеживания попыток
func checkResult(num int) (int, string) {
	// if guesses >= 6 {
	// 	return 0, errors.New("too many attempts")
	// }
	attemps++
	if num > random {
		str := fmt.Sprintln("Загаданное число меньше.")
		return -1, str
	}
	if num < random {
		str := fmt.Sprintln("Загаданное число больше.")
		return 1, str
	}
	str := fmt.Sprintf("Правильно! Вы угадали число с %d попытки.\n", attemps)
	return 0, str
}
