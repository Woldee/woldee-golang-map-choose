package main

import "fmt"

func main() {
	listFunction := 0
	arrValue := map[string]string{
		"YANDEX": "HTTPS//TEST",
	}
	for listFunction != 4 {

		listFunction = chooseValue()

		switch listFunction {
		case 1:
			fmt.Println("\nРезультат: ")
			lookPages(arrValue)
		case 2:
			keyValue, valValue := addPages()
			fmt.Println("\nРезультат: ")
			arrValue[keyValue] = valValue
		case 3:
			keyValue := deletePages()
			fmt.Println("\nРезультат: ")
			delete(arrValue, keyValue)
		}
	}
}

func chooseValue() int {
	listFunction := 0
	fmt.Println("\nВведите что вам нужно")
	fmt.Println("1 - Посмотреть закладки")
	fmt.Println("2 - Добавить закладку")
	fmt.Println("3 - Удалить закладку")
	fmt.Println("4 - Завершить работу")
	fmt.Print("-----> ")
	fmt.Scan(&listFunction)
	return listFunction

}

func lookPages(arrValue map[string]string) {
	for key, value := range arrValue {
		fmt.Println(key, value)
	}
}

func addPages() (string, string) {
	var keyValue string
	var valValue string
	fmt.Println("Введите ключ")
	fmt.Scan(&keyValue)
	fmt.Println("Введите значение")
	fmt.Scan(&valValue)
	return keyValue, valValue
}

func deletePages() string {
	var keyValue string
	fmt.Println("Введите ключ, который вы хотите удалить")
	fmt.Scan(&keyValue)
	return keyValue
}
