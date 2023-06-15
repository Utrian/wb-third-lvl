package main

type House struct {
	windowType string
	doorType   string
	floor      int
}

// Реализуем билдер - который представляет из себя
// интерфейс. Такая реализация нужна, когда
// билдеров реализующих этот интерфейс несколько.
// В обратном случае, можно обойтись структурой.

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
