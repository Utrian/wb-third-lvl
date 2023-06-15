package main

// Интерфейс объектов, над которымм мы собираемся
// проводить операции
type Shape interface {
	getType() string
	accept(Visitor)
}
