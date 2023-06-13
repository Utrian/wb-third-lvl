package main

import (
	"fmt"
	"time"
)

// The or-channel
// https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html#callout_concurrency_patterns_in_go_CO5-1

func main() {
	start := time.Now()

	<-OR(
		sig(100*time.Millisecond),
		sig(101*time.Millisecond),
		sig(500*time.Millisecond),
	)

	fmt.Printf("done after %v", time.Since(start))
}

// Возвращает сразу канал, и в фоне спит выделенное время
// после сна закрывает канал.
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		time.Sleep(after)
		defer close(c)
	}()
	return c
}

func OR(channels ...<-chan interface{}) <-chan interface{} {
	// Т.к. функция рекурсиная - опеределяем условия завершения:
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	// Создаем горутину, чтобы ожидать сообщения
	// без блокировки
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		// Кейс для оптимизации рекурсии - ограничение
		// кол-ва горутин, на случай, если каналов только 2
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-OR(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
