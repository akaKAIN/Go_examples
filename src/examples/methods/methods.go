package main

import "fmt"

type Car struct {
	price uint32
	color string
	isStart bool
}

func (car *Car) StartCar() error {
	if car.isStart {
		err := fmt.Errorf("Машина уже работает")
		return err
	}
	car.isStart = true
	fmt.Println("Машина завелась")
	return nil
}

func (car *Car) StatusCar() {
	if car.isStart {
		fmt.Println("Мишина заведена и готова")
		return
	}
	fmt.Println("Машина выключена")
}

func (car *Car) CarInfo() {
	var status string
	switch car.isStart {
	case true:
		status = "готова к тест-драйву"
	case false:
		status = "выключена"

	}
	fmt.Printf("Машина с цветом %s и стоимостью %d, сей час %s\n", car.color, car.price, status)
}

func (car Car) Testing(){
	fmt.Println("Начался ТЕСТ-ДРАЙВ")
	car.isStart = false
}

func main(){
	var myCar = &Car{32000, "Red", false}

	myCar.StatusCar()
	myCar.CarInfo()
	err := myCar.StartCar()
	if err != nil {
		fmt.Println(err)
	}
	myCar.StatusCar()
	myCar.CarInfo()
	myCar.Testing()
	myCar.CarInfo()

}
