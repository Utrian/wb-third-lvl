package main

// Интерфейс получателя
type Device interface {
	on()
	off()
}
