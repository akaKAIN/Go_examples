package main

import (
	"encoding/json"
	"examples/struct"
	"fmt"
	"image/gif"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	const fileName = "data.json"

	fmt.Printf("Запущена программа создания учетных записей. Все созданные записи сохранятся в файле в формате <json>\n\n")
Point:
	var count = new(uint32)
	fmt.Println("Введите количество создаваемых учетных записей (число):")
	if _, err := fmt.Scan(&count); err != nil {
		log.Printf("Ошибка ввода данных:\n%s", err)
		goto Point

	}
	startTime := time.Now()
	var accs = _struct.Creator(*count)
	rawData, err := json.MarshalIndent(accs, "", "  ")
	if err != nil {
		log.Fatal("Ошибка Ident:\n", err)
	}
OpenLoop:
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Файла для записи не существует:\n%s", err)
		file, err = os.Create(fileName)
		fmt.Printf("Инициализация создания файла <%s>...\n", fileName)
		goto OpenLoop
	}
	defer file.Close()
	if err = ioutil.WriteFile(fileName, rawData, 0); err == nil {
		fmt.Printf("Данные успешно сгенерированы и сохранены в файл.\n" +
			"Время выполенения  служебных процессов - ")
	}
	endTime := time.Now()
	fmt.Printf("%v\n", endTime.Sub(startTime))
	gif.Decode()
}
