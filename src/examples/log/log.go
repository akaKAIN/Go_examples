package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := os.Chmod("/home/kain/Git/Bill/Golang/GogoProgect/src/examples/log/some.txt", 0777); err != nil {
		log.Fatal(err)
	}
	dir, _ := os.Getwd()

	fmt.Printf("dir = %s\n", dir)
	fileName := "testFileName.txt"
	fullPath := dir + "/" + fileName
	if err := os.Remove(fileName); err != nil {
		log.Printf("Файл <%s> не найден\n", fullPath)
	}
	fmt.Printf("os.Getgid= %d\nos.GetPageSize= %d\n", os.Getegid(), os.Getpagesize())
	hostName, hostErr := os.Hostname()
	fmt.Printf("os.Host: %s, err: %v\n", hostName, hostErr)
}
