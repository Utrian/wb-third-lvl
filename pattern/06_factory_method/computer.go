package main

import "fmt"

// Типы, создаваемые фабричным методом
const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

// Интерфейс для создаваемых объектов
type Computer interface {
	GetType() string
	PrintDetails()
}

// Фабричный метод
func New(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s - несуществующий тип!\n", typeName)
		return nil
	case ServerType:
		return newServer()
	case PersonalComputerType:
		return newPersonalComputer()
	case NotebookType:
		return newNotebook()
	}
}
