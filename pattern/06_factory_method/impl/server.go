package main

import "fmt"

// Объявляем объект
type Server struct {
	Type   string
	Core   int
	Memory int
}

// Реализуем интерфейс Computer
func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetails() {
	fmt.Printf("%s Core:[%d] Memory: [%d]\n", s.Type, s.Core, s.Memory)
}

// Базовый конструктор
func newServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   32,
		Memory: 256,
	}
}
