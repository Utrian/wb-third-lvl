package main

import (
	"fmt"
	"log"
)

/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
В данном случае WalletFacade является основным фасадом
у которого есть методы, например, addMoneyToWallet.
Но сначала остановимся на newWalletFacade, который тоже
является фасадом, т.к. ограничивает пользователя от
реализации создания структуры - в создании участвуют
такие методы как newAccount, newWallet, но нам как
пользователю об этих методах знать не нужно - это
внутренняя реализация, а нужно нам только newWalletFacade.
По поводу addMoneyToWallet - это метод интерфейса WalletFacade,
который также имеет вызов незначимых для пользователя функций
внутри себя.
*/
type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	return nil
}

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
