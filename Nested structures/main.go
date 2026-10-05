package main

import (
	"fmt"
	"strings"
)

// User представляет пользователя
type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address Address
	Cart    []CartItem
}

// Address представляет адрес пользователя
type Address struct {
	Street     string
	City       string
	PostalCode string
}

// CartItem представляет элемент в корзине
type CartItem struct {
	Product  Product
	Quantity int
}

// Product представляет продукт в корзине
type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Brand       string
	Rating      float64
	Reviews     int
}

func main() {

	User := User{
		ID:    1,
		Name:  "Иван Петров",
		Email: "ivan.petrov@example.com",
		Phone: "+7 999 123-45-67",
		Address: Address{
			Street:     "Улица Ленина",
			City:       "Москва",
			PostalCode: "101000",
		},
		Cart: []CartItem{
			{
				Product: Product{
					ID:          1,
					Name:        "Ноутбук",
					Description: "Мощный ноутбук для работы и игр",
					Price:       59990,
					Category:    "Электроника",
					Brand:       "Brand A",
					Rating:      4.5,
					Reviews:     120,
				},
				Quantity: 1,
			},
			{
				Product: Product{
					ID:          2,
					Name:        "Смартфон",
					Description: "Современный смартфон с отличной камерой",
					Price:       29990,
					Category:    "Электроника",
					Brand:       "Brand B",
					Rating:      4.7,
					Reviews:     200,
				},
				Quantity: 2,
			},
			{
				Product: Product{
					ID:          3,
					Name:        "Наушники",
					Description: "Беспроводные наушники с шумоподавлением",
					Price:       7990,
					Category:    "Аудио",
					Brand:       "Brand C",
					Rating:      4.3,
					Reviews:     80,
				},
				Quantity: 1,
			},
		},
	}

	printInfo(User)

}

func printInfo(user User) {

	//Информация о пользователе.
	fmt.Printf("Покупатель %v. Телефон: %v. Адрес: г. %v, %v.\n", user.Name, user.Phone, user.Address.City, user.Address.Street)

	electronicByerStatus := false
	var electronicByerStatusString string
	productPriceOver1000 := []string{}
	var productPriceOver1000List string

	totalCartPrice := 0

	for _, v := range user.Cart {
		if v.Product.Category == "Электроника" {
			electronicByerStatus = true
		}
		if v.Product.Price > 10000 {
			productPriceOver1000 = append(productPriceOver1000, v.Product.Name)
		}
		totalCartPrice += v.Product.Price * v.Quantity
	}

	if electronicByerStatus {
		electronicByerStatusString = "является"
	} else {
		electronicByerStatusString = "не является"
	}

	//Покупатель электроники.
	fmt.Printf("Пользователь %s покупателем электроники.\n", electronicByerStatusString)

	if len(productPriceOver1000) > 0 {
		productPriceOver1000List = strings.Join(productPriceOver1000, ", ")
	} else {
		productPriceOver1000List = "отсутствуют"
	}
	//Товары с высокой ценой
	fmt.Printf("Товары в корзине, где цена 10000 и более: %s.\n", productPriceOver1000List)

	//Общая сумма покупки
	fmt.Printf("Общая сумма покупки: %d руб.\n", totalCartPrice)

}
