package main

import "fmt"

//Объявление интерфейса
type Payer interface {
	Pay(int) error
}

//Объявление структуры "Кошелька"
type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Недостаточно денег в кошельке.")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer){
	if err := p.Pay(10); err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T.\n\n", p)
}

func main(){
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &
}