// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		"Клавиатура JZ9": 19200
		"Наушники N45": 9600
		"Смартфон S10": 55000
	*/
	var goods1 string = "Клавиатура JZ9"
	var goods2 string = "Наушники N45"
	var goods3 string = "Смартфон S10"

	var price1 float64 = 19200
	var price2 float64 = 9600
	var price3 float64 = 55000

	var input string
	fmt.Printf("Введите название товара:")
	if _, err := fmt.Scan(&input); err != nil {
		fmt.Printf("Неверный ввод товара: %s\n", err)
	}
	//
	switch {
	case strings.Contains(strings.ToLower(goods1), strings.ToLower(input)):
		fmt.Printf("%s: %.0f\n", goods1, price1)
	case strings.Contains(strings.ToLower(goods2), strings.ToLower(input)):
		fmt.Printf("%s: %.0f\n", goods2, price2)
	case strings.Contains(strings.ToLower(goods3), strings.ToLower(input)):
		fmt.Printf("%s: %.0f\n", goods3, price3)
	default:
		fmt.Printf("Товар %s не найден.\n", input)
	}

}
