package main

import "fmt"

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	fmt.Println(areaCalculator.area) // 4

	circle.accept(areaCalculator)
	fmt.Println(areaCalculator.area) // 28.259999
}
