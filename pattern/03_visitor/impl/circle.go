package main

// Конкретная реализация фигуры
type Circle struct {
	radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}
