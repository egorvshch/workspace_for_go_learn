package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	m1 := map[string]map[string]float64{
		"Экшен": {
			"Фильм1": 8.52,
			"Фильм2": 6.0,
		},
		"Драма": {
			"Пильм3": 7.524,
			"Аильм4": 7.524,
			"Фильм5": 5.54,
		},
	}

	printRecommendations(m1)

}

func printRecommendations(movies map[string]map[string]float64) {

	tempMap := map[string]map[float64][]string{}  // временная мапа мап
	subTempMap := map[float64][]string{}          // временная вложенная мапа, обратная ключ - значение

	for key1, subMap := range movies {				
		for key2, val := range subMap {
			subTempMap[val] = append(subTempMap[val], key2) 			// меняем местами ключ - значения (рейтинг и фильмы)
			tempMap[key1] = subTempMap
		}
	}

	sliceZhanr := []string{}			// создаем слайс для упорядочивания ключей жанров

	for key := range tempMap {
		sliceZhanr = append(sliceZhanr, key)
	}
	slices.Sort(sliceZhanr)

	for _, val1 := range sliceZhanr {	// итерируемся по упорядоченному слайсу жанров

		sliceRaitings := []float64{}		// создаем слайс для упорядочивания ключей рейтинга

		for k := range tempMap[val1] {

			sliceRaitings = append(sliceRaitings, k)
			slices.Sort(sliceRaitings)

		}

		fmt.Printf("%s: ", val1)		// распечатываем жанр

		for _, val2 := range slices.Backward(sliceRaitings) {    // итерируемся по упорядоченному слайсу рейтинга в обратном порядке (от большего к меньшему)
			if val2 > 7 {
				slices.Sort(tempMap[val1][val2])			// упорядочиваем слайс фильмов
				fmt.Printf(" %s (%.2f)", strings.Join(tempMap[val1][val2], ", "), val2)   // выводим перечень фильмов и их рейтинг
			}

		}
		fmt.Println(".")

	}

}

/*
1. Функция должна принимать один аргумент movies типа map[string]map[string]float64, где:

Ключ первого уровня (string) — это название жанра (например, "Экшен", "Драма").
Ключ второго уровня (string) — это название фильма.
Значение (float64) — это рейтинг фильма.

2. Функция должна выводить на экран все жанры в алфавитном порядке, в которых есть хотя бы один фильм с рейтингом 7 и выше.
3. Для каждого жанра, в котором есть фильмы с рейтингом 7 и выше, необходимо вывести названия всех фильмов (с рейтингом 7 и выше) в порядке убывания их рейтинга, если рейтинг одинаков, тогда сортировать такие фильмы нужно в алфавитном порядке.

4. Если в жанре нет фильмов с рейтингом 7 и выше, в таком случае, жанр вовсе не должен выводиться.

Пример:
m := map[string]map[string]float64{
  "Экшен": {
    "Фильм1": 8.52,
    "Фильм2": 6.0,
  },
  "Драма": {
    "Фильм3": 7.524,
    "Фильм4": 7.527,
    "Фильм5": 5.54,
  },
}
Ожидаемый вывод:
Драма: Фильм4 (7.5), Фильм3 (7.5).
Экшен: Фильм1 (8.5).
*/
