package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0. */

func main() {
	// Объявляем три переменных: исходное число, индекс изменяемого бита и новое значение этого бита.
	var number int
	var i int
	var bit string
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	binString := fmt.Sprintf("%b", number)
	fmt.Printf("Бинарное представление введёного числа: %s\n", binString)
	fmt.Print("Введите индекс изменяемого бита: ")
	fmt.Scan(&i)
	fmt.Print("Введите новое значение бита: ")
	fmt.Scan(&bit)

	// Создаём из строки  слайс для изменения значения i-го бита
	splitString := strings.Split(binString, "")
	// Меняем значение бита по индексу в слайсе
	splitString[len(splitString)-i-1] = bit
	// Парсим изменнёный слайс в строку, а затем в целое число
	newNumber, err := strconv.ParseInt(strings.Join(splitString, ""), 2, 64)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("Битовое представление нового числа: %b",newNumber))
	fmt.Println("Новое число:",newNumber)
}
