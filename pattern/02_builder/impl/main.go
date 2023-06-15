package main

import "fmt"

func main() {
	// Создаем два билдера, реализующих один интерфейс-билдер
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	// Директор нужен для того, чтобы он сам создал
	// объект согласно какому-то внутреннему алгоритму
	// вызова методов билдера (в данном случае, этот
	// алгоритм - buildHouse).
	// Директор сам по себе является симбионтом, а не
	// неотъемлемой частью билдера.
	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)
	// Normal House Door Type: Wooden Door
	// Normal House Window Type: Wooden Window
	// Normal House Num Floor: 2

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

	// Igloo House Door Type: Snow Door
	// Igloo House Window Type: Snow Window
	// Igloo House Num Floor: 1
}
