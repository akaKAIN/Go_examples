package pathing

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
	Tree          = "|"
	Branch        = "---"
)

func WorkWithPath() {
	fmt.Println("Enter into package---------------")
	projectPath, _ := os.Getwd()
	fmt.Println("projectPath:", projectPath)

	AbsPath, err := filepath.Abs("log")
	fmt.Printf("filepath(Absath of log): %v - err: %v\n", AbsPath, err)

	//Выводит имя базового файла
	val1 := filepath.Base(AbsPath)
	fmt.Printf("%T(Base): %s\n", val1, val1)

	val2 := filepath.Dir(AbsPath)
	fmt.Printf("Dir: %v\n", val2)

	fmt.Printf("No dots: %s\n", filepath.Ext("index"))
	fmt.Printf("One dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js"))

	arr, err := filepath.Glob("src")
	fmt.Println(arr, err)

	a, b, c := "one", "two", "three"
	fmt.Println(filepath.Join(a, b, c))

	//Walk functionality
	var count int
	err = filepath.Walk("/home/kain/Git/Bill/Golang/", func(path string, _ os.FileInfo, _ error) error {
		count++
		name, err := os.Stat(path)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		var answerRow string
		switch name.IsDir() {
		case true:
			answerRow = Tree
		case false:
			answerRow = Branch
		}
		answerRow += filepath.Base(path)
		fmt.Printf("path: %s,\tBase: %s\t IsDir: %t\n", path, filepath.Base(path), name.IsDir())
		//fmt.Printf("%v\n", answerRow)
		return nil
	})
	fmt.Println(count, time.Now())

}

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() && f.Name() == ".git" {
		return filepath.SkipDir
	}
	fmt.Printf("Visited: %s\n", path)
	return nil
}
