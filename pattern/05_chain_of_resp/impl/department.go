package main

// Интерфейс обработчика

type Department interface {
	execute(*Patient)
	setNext(Department)
}
