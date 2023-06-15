package main

import (
	"fmt"
)

// Реализуем структуру, реализующую интерфейс
// посетителя
type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
	a.area = s.side * s.side
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
	a.area = 3.14 * c.radius * c.radius
}
