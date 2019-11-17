package main

import (
	"fmt"
	"sort"
	"sync"
)

type MySum struct {
	SumMutex *sync.Mutex
	Sum      []int

}

func main() {
	wg := new(sync.WaitGroup)
	obj := MySum{
		Sum:      make([]int, 0),
		SumMutex: new(sync.Mutex),
	}
	for i := 1; i < 100; i++ {
		go sorting(i, &obj, wg)
	}
	//time.Sleep(2)
	for ind, elem := range obj.Sum {
		fmt.Println(ind, elem)
	}
	fmt.Println(len(obj.Sum))
	wg.Wait()
	sort.Ints(obj.Sum)
	for _, elem := range obj.Sum {
		fmt.Println(elem)
	}
}

func sorting(n int, obj *MySum, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	if n%2 == 0 {
		go even(n, obj)
	} else {
		go odd(n, obj)
	}
}

func even(n int, obj *MySum) {
	obj.SumMutex.Lock()
	defer obj.SumMutex.Unlock()
	obj.Sum = append(obj.Sum, n)

}
func odd(n int, obj *MySum) {
	obj.SumMutex.Lock()
	defer obj.SumMutex.Unlock()
	obj.Sum = append(obj.Sum, n)
}
