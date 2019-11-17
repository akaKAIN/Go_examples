package main

import (
	"fmt"
	"log"
)

type EcoInt uint32

func main() {
	var i interface{}
	var inc, cost, input EcoInt
	text, err := fmt.Scan(&input)
	if err != nil {
		log.Println("errorroror-", err)
	}
	fmt.Printf("text=%v, type=%T\n%v\n\n", input, input, text)
	//numText, err := strconv.ParseInt(text, 10,32)
	//if err != nil {
	//	fmt.Println("Error", err)
	//}
	//fmt.Println(numText)
	//err :=  (&cost)
	i = cost
	s := i.(EcoInt) + 1
	fmt.Println("s=", s)
	for inc =0 ; inc < cost; inc ++ {

	}
}
