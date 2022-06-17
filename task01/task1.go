package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).*/

type Human struct {
	age int
}

func (h Human) getAge() int {
	return h.age
}

type Action struct {
	Human
}

/* Метод можно переопределять
func (a Action) getAge() int {
return a.age + 10
}
a.getAge() : 35
a.Human.getAge() : 25 */

func main() {
	//создаём структуру Human
	h := Human{25}
	//создаём структуру Action, в которую встроена структура Human
	a := Action{h}
	//Вызываем у структуры Action метод getAge структуры Human
	fmt.Println(a.getAge())
	fmt.Println(a.age)
}
