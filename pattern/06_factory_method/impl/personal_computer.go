package main

import "fmt"

// Объявляем объект
type PC struct {
	Type    string
	Core    int
	Memory  int
	Display string
}

// Реализуем интерфейс Computer
func (pc PC) GetType() string {
	return pc.Type
}

func (pc PC) PrintDetails() {
	fmt.Printf("%s Core:[%d] Memory: [%d] Display: [%s]\n", pc.Type, pc.Core, pc.Memory, pc.Display)
}

// Базовый конструктор
func newPersonalComputer() Computer {
	return PC{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Display: "2560x1440",
	}
}
