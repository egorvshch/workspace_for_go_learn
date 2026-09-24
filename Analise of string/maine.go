package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if str, err := GetInput(); err != nil {
		fmt.Printf("Error cause %q\n", err)
		os.Exit(1)
	} else {
		letters, digits, spaces, punctuation := CountCharacters(str)
		DisplayResults(letters, digits, spaces, punctuation)
	}
}

// Ввод текста
func GetInput() (string, error) {

	for {
		fmt.Print("Введите текст: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		textInput := scanner.Text()
		if err != nil {
			return "", fmt.Errorf("Произошла ошибка ввода %v", err)
		}
		if len(strings.TrimSpace(textInput)) == 0 {
			fmt.Println("Строка не должна быть пустой либо содержать только пробелы")
		} else {
			return textInput, nil
		}

	}
}

// Подсчет символов
func CountCharacters(text string) (letters, digits, spaces, punctuation int) {

	for _, char := range text {
		if char == ' ' { // пробелов (spaces)
			spaces++
		} else if char == '.' || // знаков препинания (punctuation)
			char == ',' ||
			char == '?' ||
			char == '!' ||
			char == ':' ||
			char == ';' ||
			char == '-' ||
			char == '"' ||
			char == '(' ||
			char == ')' {
			punctuation++
		} else if char == '0' || // цифр (digits)
			char == '1' ||
			char == '2' ||
			char == '3' ||
			char == '4' ||
			char == '5' ||
			char == '6' ||
			char == '7' ||
			char == '8' ||
			char == '9' {
			digits++
		} else {
			letters++ // букв (letters)
		}
	}
	return letters, digits, spaces, punctuation

}

// Вывод результатов
func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Printf("Количество букв: %d\n", letters)
	fmt.Printf("Количество цифр: %d\n", digits)
	fmt.Printf("Количество пробелов: %d\n", spaces)
	fmt.Printf("Количество знаков препинания: %d\n", punctuation)

}
