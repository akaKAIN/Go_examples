package main

import "fmt"

func main() {
	test()
}

func test(){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			if err, ok := r.(error); ok {
				fmt.Println(err)
			}
		}
	}()
	panic("Errrrrrr")

}
