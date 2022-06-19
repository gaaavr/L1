package main

import (
	"context"
	"fmt"
	"time"
)

/* Реализовать все возможные способы остановки выполнения горутины.  */


func main() {
	// Горутина завершает работу по сигналу в специальный канал
	die := make(chan struct{})
	go goOne(die)
	close(die)

	// Горутина завершает работу при отмене контекста
	ctx, cancel := context.WithCancel(context.Background())
	go goTwo(ctx)
	cancel()

	// Горутина завершает работу при достижении какого либо времени с помощью контекста
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond))
	go goThree(ctx)
	defer cancel()

	// Горутина завершает работу через указанный промежуток времени
	go goFour()
	time.Sleep(4 * time.Second)
}

func goOne(end chan struct{}) {
	for {
		select {
		case <-end:
			fmt.Println("Stop 1")
			return
		}
	}
}

func goTwo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop 2")
			return
		}
	}
}

func goThree(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop 3")
			return
		}
	}
}

func goFour() {
	t := time.After(time.Second*2)
	for {
		select {
		case <-t:
			fmt.Println("Stop 4")
			return
		}
	}
}
