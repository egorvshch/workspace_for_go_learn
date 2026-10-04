package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	friendsData := map[string][]string{
		"Алексей":  {"Иван", "Сергей", "Елена"},
		"Иван":     {"Алексей", "Дмитрий", "Мария"},
		"Сергей":   {"Алексей", "Елена"},
		"Дмитрий":  {"Иван", "Елена", "Ольга"},
		"Елена":    {"Алексей", "Сергей", "Дмитрий"},
		"Мария":    {"Иван", "Ольга"},
		"Ольга":    {"Дмитрий", "Мария"},
		"Анна":     {"Петр"},
		"Петр":     {"Анна", "Сергей"},
		"Светлана": {"Иван", "Елена"},
	}

	//Подсчет друзей
	countFriends := countFriends(friendsData)
	fmt.Printf("Количество друзей:\n")
	for user, count := range countFriends {
		fmt.Printf("%s: %d\n", user, count)
	}

	//Общие друзья
	user1 := "Иван"
	user2 := "Елена"
	FiendList := commonFriends(friendsData, user1, user2)
	fmt.Printf("Общие друзья между пользователями %s и %s: %s.\n", user1, user2, strings.Join(FiendList, ", "))

	//Наиболее популярные пользователи
	maxFriendUsers, countFriend := mostPopularUsers(friendsData)
	fmt.Printf("Наиболее популярные пользователи: %s (количество друзей: %d).", strings.Join(maxFriendUsers, ", "), countFriend)

}

func countFriends(users map[string][]string) map[string]int {

	/*
		Реализуйте функцию countFriends, которая принимает map с пользователями и их друзьями и возвращает map, где ключом является имя пользователя, а значением — количество его друзей.
	*/

	counterFriend := map[string]int{}

	for user, friend := range users {
		counterFriend[user] = len(friend)
	}
	return counterFriend
}

func commonFriends(users map[string][]string, user1, user2 string) []string {

	/*
		Реализуйте функцию commonFriends, которая принимает map с пользователями и их друзьями, а также имена двух пользователей. Функция должна возвращать список общих друзей между этими двумя пользователями.
	*/

	commonFiendList := []string{}

	for _, friend := range users[user1] {
		if slices.Contains(users[user2], friend) {
			commonFiendList = append(commonFiendList, friend)
		}

	}
	return commonFiendList
}

func mostPopularUsers(users map[string][]string) ([]string, int) {

	/*
		Реализуйте функцию mostPopularUsers, которая принимает map с пользователями и их друзьями и возвращает список имен пользователей с наибольшим количеством друзей и это количество.
	*/

	counterFriend := map[string]int{}
	maxFriends := 0
	maxFriendsUsers := []string{}
	for user, friend := range users {
		counterFriend[user] = len(friend)
		if len(friend) > maxFriends {
			maxFriends = len(friend)
		}
	}

	for user := range counterFriend {
		if counterFriend[user] == maxFriends {
			maxFriendsUsers = append(maxFriendsUsers, user)
		}
	}
	slices.Sort(maxFriendsUsers)
	return maxFriendsUsers, maxFriends

}
