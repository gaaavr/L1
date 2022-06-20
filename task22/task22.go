package main

import (
	"fmt"
	"math/big"
)

/* Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.*/

func main() {
	// Объявляем переменные и вводим их значения
	var (
		a, b float64
		s    string
	)
	fmt.Println("Введите первое и второе число:")
	fmt.Scan(&a, &b)
	// Пакет big реализует операции с большими числами
	// Переводим введёные числа в "большие числа"
	bigA := big.NewFloat(a)
	bigB := big.NewFloat(b)
	fmt.Println("Введите операцию:")
	fmt.Scan(&s)
	// Проверяем, какая операция была введена и выполняем её
	switch s {
	case "+":
		sum := new(big.Float)
		sum = sum.Add(bigA, bigB)
		fmt.Println(sum.String())
	case "-":
		sub := new(big.Float)
		sub = sub.Sub(bigA, bigB)
		fmt.Println(sub.String())
	case "*":
		mul := new(big.Float)
		mul = mul.Mul(bigA, bigB)
		fmt.Println(mul.String())
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
		} else {
			div := new(big.Float)
			div = div.Quo(bigA, bigB)
			fmt.Println(div.String())
		}
	default:
		fmt.Println("invalid operation")
	}

}
