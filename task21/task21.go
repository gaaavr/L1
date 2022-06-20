package main

import "fmt"

/* Реализовать паттерн «адаптер» на любом примере.*/

// Создадим интерфейс, описывающий вставку адаптера в квадратный порт

type computer interface {
	insertInSquarePort()
}

// Создадим структуру mac, у которой есть метод, удовлетворяющий интерфейсу выше
type mac struct {
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Квадратный адаптер вставлен в порт Mac")
}

// Создадим структуру windows, которая не удовлетворяет интерфейсу выше, т.к.
// у машины на винде круглый порт

type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Круглый адаптер вставлен в порт машины на Windows")
}

// Создадим структуру адаптер, методы которой удовлетворяю интерфейсу выше.
// Вставим в в неё нашу структуру windows.

type windowsAdapter struct {
	windowMachine *windows
}

// Метод адаптера удовлетворяет интерфейсу,
// внутри метода вызывается метод windows и всё благополучно работает

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

// Также есть структура client, метод которой принимает на вход интерфейс и вызывает его метод
type client struct {
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}

func main() {
	// Создадим экземпляры машины на винде и мака
	mac := &mac{}
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	// При вызове одинаковых методов у разных машин выведутся
	// результаты в зависимости от машины. Т.к. для windows существует "адаптер".
	mac.insertInSquarePort()
	windowsMachineAdapter.insertInSquarePort()
	// Создадим экземпляр клиента.
	client := &client{}
	// При вызове метода клиента сработает метод для mac. Аналогично будет для windows.
	client.insertSquareUsbInComputer(mac)
	client.insertSquareUsbInComputer(windowsMachineAdapter)

}