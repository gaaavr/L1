package main

import (
	"L1/task24/point"
	"fmt"
)

/* Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с
инкапсулированными параметрами x,y и конструктором.*/

func main() {
	// Создаём две точки с помощью конструктора из пакета point
	pointOne := point.NewPoint(2, 5)
	pointTwo := point.NewPoint(5, 8)
	// Выводим вычисленную дистанцию между точками
	fmt.Println(point.Distance(pointOne, pointTwo))
}
