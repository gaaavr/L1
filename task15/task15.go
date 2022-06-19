package main

import "fmt"

/* К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.
 */

var justString string

func createHugeString(n int) string {
	m := make([]rune, n, n)
	for i := 0; i < n; i++ {
		m[i] = 'ы'
	}
	return string(m)
}

func someFunc() {
	v := createHugeString(1 << 10) // 1 << 10 = 2 ^ 10 = 1024 символов 'ы'
	fmt.Println(v)
	// При обращении к строке по индексу, мы обращаемся к байтам, т.к. строука - это срез байт под капотом.
	// Кириллица занимает 2 байта, при таком отрезке строки будет выведено 50 символов кириллицы, а не 100.
	justString = v[:100]
	fmt.Println(justString)

}
func myFunc()  {
	// Если строку привести к слайсу рун, то получим ожидаемый результат
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
	myFunc()
}