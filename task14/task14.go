package main

import "fmt"

/* Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}. */

func main()  {
	// Инициализируем массив с перменными различного типа
	arr := [5]interface{}{1, 3.14, false, make(chan interface{}), "string"}
	// Применяем к нему функцию, которая в рантайме определяет тип переменной
	for _,v:=range arr{
		definitionValue(v)
	}

}

// Функция для определения типа переменной
func definitionValue(i interface{})  {
	switch i.(type) {
	case bool:
		fmt.Println("Значение переменной", i, "- bool.")
	case int:
		fmt.Println("Значение переменной", i, "- int.")
	case string:
		fmt.Println("Значение переменной", i, "- string.")
	case chan interface{}:
		fmt.Println("Значение переменной", i, "- chan.")
	default:
		fmt.Println("Тип не определён.")

	}
}