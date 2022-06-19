package main

import (
	"fmt"
	"strconv"
	"sync"
)

/* Реализовать конкурентную запись данных в map.  */


// Создадим структуру, полями которой будут мапа и мьютекс
// RWMutex помогает реализовать конкурентную запись и чтение в map
type concMap struct {
	data map[string]string
	mu   sync.RWMutex
}

// Реализуем два метода для добавления данных в мапу и их чтения

// При добавлении данных  в мапу одновременно писать их может только одна горутина

func (m *concMap) Add(key, value string) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}


// При получении данных из мапы одновременно читать их могут несколько горутин

func (m concMap) Get(key string) (string, bool) {
	m.mu.RLock()
	value, ok := m.data[key]
	m.mu.RUnlock()
	return value, ok
}

func main() {
	// Создали экземпляр структуры мапы с конкурентной записью
	m := concMap{
		data: make(map[string]string),
		mu:   sync.RWMutex{},
	}
	// Обьявили вэйтгруппу для ожидания главной горутиной окончания работы остальных горутин
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			m.Add("key"+strconv.Itoa(i), "value"+strconv.Itoa(i*10))
			defer wg.Done()
		}(&wg, i)
	}
	// Ждём завершения работы горутин записи в мапу
	wg.Wait()
	// Выводим записанныев мапу значения в консоль
	for i := 0; i < len(m.data); i++ {
		value, _ := m.Get("key" + strconv.Itoa(i))
		fmt.Println(value)
	}


}
