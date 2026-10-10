package main

import (
	"errors"
	"fmt"
)

type TagManager struct {
	Tags map[string]struct{}
}

func (t *TagManager) AddTag(tag string) error {
	//добавляет новый тег. Если тег уже существует, метод вернет ошибку
	if _, ok := t.Tags[tag]; ok {
		return errors.New("Ошибка, тег уже существует")
	}
	t.Tags[tag] = struct{}{}
	return nil

}

func (t *TagManager) RemoveTag(tag string) error {
	//удаляет тег по его имени. Если тега не существует, метод вернет ошибку
	if _, ok := t.Tags[tag]; !ok {
		return errors.New("Ошибка, тег не существует")
	}
	delete(t.Tags, tag)
	return nil
}

func (t *TagManager) TagExists(tag string) bool {
	//проверяtn, существует ли тег в системе
	if _, ok := t.Tags[tag]; ok {
		return true
	}
	return false
}

func (t *TagManager) ListTags() []string {
	//возвращает все уникальные теги в системе.
	sliceResult := []string{}
	for key := range t.Tags {
		sliceResult = append(sliceResult, key)
	}
	return sliceResult
}

func NewTagManager() *TagManager {
	return &TagManager{
		Tags: map[string]struct{}{},
	}
}

func main() {
	tm := NewTagManager()

	// Добавление тегов
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err)
	}

	if err := tm.AddTag("programming"); err != nil {
		fmt.Println(err)
	}

	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег уже существует
	}

	// Проверка существования тегов
	fmt.Println("Тег 'golang' существует:", tm.TagExists("golang")) // true
	fmt.Println("Тег 'python' существует:", tm.TagExists("python")) // false

	// Список тегов
	fmt.Println("Current tags:", tm.ListTags()) // [golang programming]

	// Удаление тегов
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err)
	}

	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег не существует
	}

	// Список тегов после удаления
	fmt.Println("Current tags after removal:", tm.ListTags()) // [programming]
}
