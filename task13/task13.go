package main

import "fmt"

/* Поменять местами два числа без создания временной переменной. */

func main()  {
	// Заводим две целочисленные переменные
	var (
		a = 4
		b=8
	)
	fmt.Println(a,b)
	// Меняем их значения местами
	a,b=b,a
	fmt.Println(a,b)

}