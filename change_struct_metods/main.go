package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"time"
	"unicode"
)

type User struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

// генерирует секретное имя. Составляется из первых букв имени и фамилии, за которыми следует случайное число от 1 до 100.
// Например, для "Алексей Смирнов" мы можем получить строку "АС42".
func (u *User) SecretIdentity() string {
	random := rand.IntN(100) + 1
	runesFirstName := []rune(u.FirstName)
	runesLastName := []rune(u.LastName)
	var firstCharFirstName string
	var firstCharLastName string
	if len(runesFirstName) > 0 {
		firstCharFirstName = string(runesFirstName[0])
	} else {
		return "FirstName is empty"
	}
	if len(runesLastName) > 0 {
		firstCharLastName = string(runesLastName[0])
	} else {
		return "LastName is empty"
	}
	return fmt.Sprintf("%s%s%d", firstCharFirstName, firstCharLastName, random)
}

// возвращает текущий возраст пользователя. Текущий год можно получить с помощью time.Now().Year()
func (u *User) Age() int {

	if (time.Now().Year() - u.BirthYear) > 0 {
		return time.Now().Year() - u.BirthYear
	}
	return 0
}

//добавляет язык в FavoriteLanguages.
/*
Если переданный язык - пустая строка, необходимо вернуть ошибку с текстом 'empty language name'.
Если переданный язык уже есть в слайсе, необходимо  вернуть ошибку с текстом 'duplicate'.
*/
func (u *User) AddFavoriteLanguage(language string) error {

	if len(language) == 0 {
		return fmt.Errorf("empty language name")
	}
	if slices.Contains(u.FavoriteLanguages, language) {
		return fmt.Errorf("duplicate")
	}
	u.FavoriteLanguages = append(u.FavoriteLanguages, language)
	return nil

}

// удаляет язык из FavoriteLanguages. Если языка нет в слайсе, необходимо вернуть ошибку с текстом 'not found'
func (u *User) RemoveFavoriteLanguage(language string) error {

	if slices.Contains(u.FavoriteLanguages, language) {
		u.FavoriteLanguages = slices.DeleteFunc(u.FavoriteLanguages, func(x string) bool {
			if x == language {
				return true
			}
			return false
		})
	} else {
		return fmt.Errorf("not found")
	}
	return nil
}

// проверяет, есть ли язык в списке любимых языков.
func (u *User) IsProgrammingLanguageFavorite(language string) bool {

	if slices.Contains(u.FavoriteLanguages, language) {
		return true
	}
	return false
}

// возвращает случайный любимый язык. Если список пуст, необходимо вернуть ошибку с текстом 'no options'
func (u *User) RandomFavoriteLanguage() (string, error) {

	if len(u.FavoriteLanguages) == 0 {
		return "", fmt.Errorf("no options")
	}
	randomIndex := rand.IntN(len(u.FavoriteLanguages))
	return u.FavoriteLanguages[randomIndex], nil
}

//возвращает строку с полным профилем пользователя в формате:
/*
Имя: Павел.
Фамилия: Тарасов.
Возраст: 35.
Список любимых языков программирования: [Go, Python].
*/
func (u *User) GenerateProfile() string {
	return fmt.Sprintf("Имя: %s.\nФамилия: %s.\nВозраст: %d.\nСписок любимых языков программирования: [%s].", u.FirstName, u.LastName, u.Age(), strings.Join(u.FavoriteLanguages, ", "))
}

//обновляет имя и фамилию пользователя.
/*
Если имя или фамилия пустые, необходимо вернуть ошибку с текстом 'empty data'.
Если имя или фамилия начинаются с маленькой буквы, необходимо вернуть ошибку с текстом 'invalid data'.
В случае любой ошибки данного метода, поля структуры не должны быть изменены.
*/
func (u *User) UpdateName(firstName, lastName string) error {

	if len(firstName) == 0 || len(lastName) == 0 {
		return fmt.Errorf("empty data")
	}

	runesFirstName := []rune(firstName)
	runesLastName := []rune(lastName)

	if unicode.ToLower(runesLastName[0]) == runesFirstName[0] || unicode.ToLower(runesLastName[0]) == runesLastName[0] {
		return fmt.Errorf("invalid data")
	} else {
		u.FirstName = firstName
		u.LastName = lastName
	}
	return nil
}

func main() {

	user1 := User{
		FirstName:         "Denis",
		LastName:          "Belov",
		BirthYear:         1901,
		FavoriteLanguages: []string{"Go", "SQL"},
	}

	fmt.Println(user1.GenerateProfile())
	fmt.Println(user1.UpdateName("Вася", "Пупкин"))
	fmt.Println(user1.GenerateProfile())
}
