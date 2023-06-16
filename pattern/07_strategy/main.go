package main

import "fmt"

/*
Пример на платёжке: есть интерфейс Payment с
методом Pay(), и есть различные стратегии/структуры -
платёжные системы, реализующие этот интерфейс
(qiwiPayment, cardPayment...). В месте где нужно
провести оплату будет код следующего вида.
*/

// Интерфейс
type Payment interface {
	Pay()
}

// Одна из стратегий
type cardPayment struct {
	cardNumber, cvcCode string
}

// Реализация интерфейса
func (cp *cardPayment) Pay() {
	fmt.Println("CARD PAY!")
}

// Конструктор
func NewCardPayment(number, cvc string) Payment {
	return &cardPayment{
		cardNumber: number,
		cvcCode:    cvc,
	}
}

// Место опеределения
func main() {
	p := 1
	numCard := "some number"
	cvcCode := "some cvc"

	var payment Payment // определяем интерфейс

	// Выбор стратегии
	switch p {
	case 1:
		payment = NewCardPayment(numCard, cvcCode)
	case 2:
		// ...
	}

	// Вне зависимости от стратегии, используем общий интерфейс
	payment.Pay() // CARD PAY!
}
