package _struct

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	firstNameList = [] string{"Ivan", "Goerge", "Bred", "Look", "Sid"}
	lastNameList  = [] string{"Ivanov", "Trebiany", "Lee", "Mikl", "Donavan"}
)

var idPerson, idAccount uint32 = 1, 1

type Person struct {
	Id   uint32    `json:"id"`
	Name string `json:"name"`
	Age  uint8    `json:"age"`
}

type Account struct {
	Id             uint32                 `json:"id"`
	Login          string              `json:"login"`
	Owner          Person              `json:"owner"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreatePerson(incomingName string, incomingAge uint8) Person {
	var newPerson = Person{
		Id:   idPerson,
		Name: incomingName,
		Age:  incomingAge,
	}
	idPerson++
	return newPerson
}

func CreateAccount(person Person) Account {
	var login = fmt.Sprintf("%v-%v", person.Name, person.Age)
	var newAccount = Account{
		Id:             idAccount,
		Login:          login,
		Owner:          person,
	}
	idAccount++
	return newAccount
}

func generatePersonList(count uint32) []Person {
	//Минимальный-максимальный возраст для структуры Person- (12 - 80)
	var minAge, maxAge = 12, 80
	var personStorage []Person
	var i uint32
	for i = 0; i < count; i++ {
		fullName := lastNameList[rand.Intn(len(lastNameList))] + " " + firstNameList[rand.Intn(len(firstNameList))]
		age := uint8(minAge + rand.Intn(maxAge-minAge+1))
		personStorage = append(personStorage, CreatePerson(fullName, age))
	}
	return personStorage
}

func generateAccountList(personList []Person) []Account {
	var accountStorage [] Account
	for _, person := range personList {
		accountStorage = append(accountStorage, CreateAccount(person))
	}
	return accountStorage
}

func Creator(personCount uint32) []Account {
	persons := generatePersonList(personCount)
	accounts := generateAccountList(persons)

	return accounts
}