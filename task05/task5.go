package main

import (
	"context"
	"fmt"
	"time"
)


/* Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться. */

func main()  {
	// инициализируем канал, через который будут проходить данные
	ch:=make(chan int)
	//объявляем и инициализируем в переменной время жизни программы
	var n int
	fmt.Println("Сколько секунд будет работать программа?")
	fmt.Scan(&n)
	//обозначаем время начала работы программы
	start:=time.Now()
	//создаём контекст с таймаутом, по истечении указанного времени работа программы завершится
	ctx,_:=context.WithTimeout(context.Background(),time.Second*time.Duration(n))
	//применяем  функцию rwrChan для записи и чтения данных через канал
	rwrChan(ctx,ch)
	//высчитываем время работы функции от начала работы до её прекращения
	dur := time.Since(start)
	//выводим время работы программы
	fmt.Println("Работа завершена, программа работала", int(dur.Seconds()), "секунд.")
}

//функция rwrChan принимает на вход созданный нами контекст и канал. По истечении указанного нами времени
//в контексте функция прекратит свою работу, запись и чтения из канала прекратятся.
func rwrChan(ctx context.Context, ch chan int) {
	//в отдельной горутине пишем данные в канал
	go func() {
		//в цикле на каждой итерации пишем новые данные в канал до тех пор, пока не истечет указанный нами таймаут.
		for i:=1; i>0; i++ {
			//как только таймаут истечет, контекст вернёт нам закрытый канал.
			//В этом случае закрываем канал и завершаем программу.
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch<-i
			}
		}
	}()
	// в главной горутине читаем данные из канала, пока канал не закроется
	for  {
		v,ok:=<-ch
		if !ok {
			return
		}
		fmt.Println(v)
	}

}