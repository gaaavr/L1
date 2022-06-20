package main

import "fmt"

/* Удалить i-ый элемент из слайса.*/

func main() {
	// Задаём исходный слайс интов
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// Задаём индекс удаляемого элемента
	var index int
	fmt.Println("Введите индекс элемента, который необходимо удалить:")
	fmt.Scan(&index)
	// Проверяем на наличие индекса в слайсе
	if index >= len(arr) || index < 0 {
		fmt.Println("invalid index")
		return
	}
	// Удаляем элемент, если он существует по введёному индексу
	arr = append(arr[:index], arr[index+1:]...)
	// Выводим результат после удаления
	fmt.Println(arr)
}
