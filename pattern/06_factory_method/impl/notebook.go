package main

import "fmt"

// Объявляем объект
type Notebook struct {
	Type         string
	Core         int
	Memory       int
	Display      string
	PowerAdapter string
}

// Реализуем интерфейс Computer
func (n Notebook) GetType() string {
	return n.Type
}

func (n Notebook) PrintDetails() {
	fmt.Printf("%s Core:[%d] Memory: [%d] Display: [%s] Power adapter: [%s]\n", n.Type, n.Core, n.Memory, n.Display, n.PowerAdapter)
}

// Базовый конструктор
func newNotebook() Computer {
	return Notebook{
		Type:         NotebookType,
		Core:         8,
		Memory:       16,
		Display:      "1920x1080",
		PowerAdapter: "320W",
	}
}
