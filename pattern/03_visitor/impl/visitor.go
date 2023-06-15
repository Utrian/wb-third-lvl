package main

// Интерфейс посетителей
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
}
