package main

import "fmt"

/* Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество. */

func main()  {
	// Исходная последовательность в виде слайс строк
	arr:=[]string{"cat", "cat", "dog", "cat", "tree"}
	// Множество в виде мапы
	m:=make(map[string]int)
	// В цикле по слайсу заполняем множество. Ключом является строка,
	// значением - количество этой строки в последовательности.
	for _,v:=range arr{
		m[v]++
	}
	// Выводим полученное множество
	fmt.Println(m)

}