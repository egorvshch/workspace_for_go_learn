package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const (
	MinPasswordLength = 4
	MinPasswordsCount = 1
	MaxPasswordsCount = 50
)

var (
	ErrPasswordLengthTooLow = errors.New("password length too low")
	ErrPasswordsCountTooLow = errors.New("too low passwords count")
	ErrPasswordsCountTooBig = errors.New("too big passwords count")
)

var (
	upperChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerChars   = []rune("abcdefghijklmnopqrstuvwxyz")
	digitChars   = []rune("0123456789")
	specialChars = []rune("!@#$%^&*")
)

// generatePassword генерирует count паролей длиной length.
func generatePassword(length int, count int) ([]string, error) {

	if length < MinPasswordLength {
		return nil, ErrPasswordLengthTooLow
	}
	if count < MinPasswordsCount {
		return nil, ErrPasswordsCountTooLow
	}
    	if count > MaxPasswordsCount {
		return nil, ErrPasswordsCountTooBig
	}

	// проверка дубликата паролей
	mapCheckPassword := make(map[string]struct{})

	// объединенный набор символов для генерации паролей
	sliceSimbolForPassword := make([]rune, 0)
	sliceSimbolForPassword = append(sliceSimbolForPassword, upperChars...)
	sliceSimbolForPassword = append(sliceSimbolForPassword, lowerChars...)
	sliceSimbolForPassword = append(sliceSimbolForPassword, digitChars...)
	sliceSimbolForPassword = append(sliceSimbolForPassword, specialChars...)

	// итоговый слайс паролей
	passwordCollection := make([]string, 0)

	for i := 0; i < count; i++ {

		// формируем символы для пароля
		slicePasswordSimbols := make([]string, 0)

		//добавляем первые 4 рандомных символа
		upperCharsKey, _ := rand.Int(rand.Reader, big.NewInt(int64(len(upperChars))))
		lowerCharsKey, _ := rand.Int(rand.Reader, big.NewInt(int64(len(lowerChars))))
		digitCharsKey, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digitChars))))
		specialCharsKey, _ := rand.Int(rand.Reader, big.NewInt(int64(len(specialChars))))

		slicePasswordSimbols = append(slicePasswordSimbols, string(upperChars[int(upperCharsKey.Int64())]))
		slicePasswordSimbols = append(slicePasswordSimbols, string(lowerChars[int(lowerCharsKey.Int64())]))
		slicePasswordSimbols = append(slicePasswordSimbols, string(digitChars[int(digitCharsKey.Int64())]))
		slicePasswordSimbols = append(slicePasswordSimbols, string(specialChars[int(specialCharsKey.Int64())]))

		// добавляем остальные символы в пароль

		for i := 0; i < length-4; i++ {
			sliceSimbolForPasswordKey, _ := rand.Int(rand.Reader, big.NewInt(int64(len(sliceSimbolForPassword))))
			slicePasswordSimbols = append(slicePasswordSimbols, string(sliceSimbolForPassword[int(sliceSimbolForPasswordKey.Int64())]))
		}
		// fmt.Println(slicePasswordSimbols)

		// дополнительно перемешиваем символы в пароле
		for i := len(slicePasswordSimbols) - 1; i >= 0; i-- {
			k, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
			j := slicePasswordSimbols[int(k.Int64())]
			slicePasswordSimbols[int(k.Int64())] = slicePasswordSimbols[i]
			slicePasswordSimbols[i] = j
		}
		// формируем строку пароля
		stringPassword := strings.Join(slicePasswordSimbols, "")

		//проверяем повторяемость пароля
		if _, ok := mapCheckPassword[stringPassword]; ok {
			count++
		}

		//добавляем пароль в итоговый слайс паролей
		mapCheckPassword[stringPassword] = struct{}{}
		passwordCollection = append(passwordCollection, stringPassword)

	}
	// fmt.Println(sliceSimbolForPassword)
	// fmt.Println(passwordCollection)

	return passwordCollection, nil
}

func main() {

	result, err := generatePassword(2, 10)
	if err != nil {
		log.Fatal("Ошибка: ", err)
	}
	fmt.Printf("Сгенерированные пароли: %v", result)
}
