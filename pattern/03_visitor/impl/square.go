package main

// Конкретная реализация фигуры
type Square struct {
	side float64
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}
