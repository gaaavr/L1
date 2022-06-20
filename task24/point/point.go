package point

import "math"

// Создаём структуру, поля которой являются координатами точки

type Point struct {
	x, y float64
}

// Создаём функцию-конструктор, которая возвращает нам готовую структуру с координатами точки

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Функция для вычисления расстояния между двумя точками

func Distance(a, b Point) float64 {
	return math.Sqrt(math.Pow((a.x-b.x), 2) + math.Pow((a.y-b.y), 2))
}
