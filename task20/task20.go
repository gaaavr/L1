package main

import (
	"fmt"
	"strings"
)

/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

func main() {
	// Применяем написаннную функцию для переворачивания слов
	reverseSentence("snow dog sun")
}

func reverseSentence(s string) {
	// Переводим исходную строку в слайс строк, в качестве разделителя используется пробел
	arr := strings.Split(s, " ")
	// Итерируемся до середины слайса
	for i := 0; i < len(arr)/2; i++ {
		// Меняем местами строки до середины слайса и после
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	// Объединяем строки в полученном слайсе в одну строку
	s = strings.Join(arr, " ")
	// Выводим результат
	fmt.Println(s)
}
