package main

import (
	"fmt"
	"log"
)

/*
Пример на торговом автомате, выдающем один товар. Автомат
быть в одном из состояний: hasItem, noItem, itemRequested,
hasMoney. И есть четыре операции: выбрать предмет, добавить
предмет, ввести деньги, выдать предмет. В начале автомат
в состоянии itemRequested. После внесения денег, он переходит
в состояние hasMoney.
В зависимости от состояния, автомат отвечает по разному на
один и тот же запрос: если мы хотим получить предмет, то
автомат удовлетворит это желание, если он находится в
состоянии hasItemState, или отклонит запрос, если состояние
noItemState.
*/

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem() // Item request
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10) // Money entered is ok
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem() // Dispensing Item
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2) // Adding 2 items
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem() // Item request
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10) // Money entered is ok
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem() // Dispensing Item
	if err != nil {
		log.Fatalf(err.Error())
	}
}
