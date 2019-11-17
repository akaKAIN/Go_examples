package main

import (
	"fmt"
	"log"
)

type Payer interface {
	Pay(uint32) error
	Raplenishment(uint32) error
}
type Wallet struct {
	Cash uint32
}


func (w *Wallet) Pay(amount uint32) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

func (w *Wallet) Replenishment(amount uint32) error {
	point := w.Cash
	fmt.Printf("Befor: point=%d, w.Cash=%d", point, w.Cash)
	w.Cash += amount
	fmt.Printf("After: point=%d, w.Cash=%d", point, w.Cash)
	if point == w.Cash {
		return fmt.Errorf("Ошибка зачисления")
	}
	return nil
}

func Buy(p Payer){
	if err := p.Pay(95); err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func Repl(p Payer){
	if err := p.Raplenishment(40); err != nil {
		panic(err)
	}
	fmt.Print("Money in ...")
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
	log.Printf("В кошельке oсталось %d", myWallet.Cash)
	Repl(myWallet)
}