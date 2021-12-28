package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Struct for representation of total slice.
// First level of JSON object parsing.
type Users struct {
	Users []User `json:"users"`
}

// Internal User representation.
// Second level of JSON object parsing.
type User struct {
	Name   string `json:"name"` // name - имя должно полностью совпадать с именем поля в json, включая регистр
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

// Social block representation.
// Third level of JSON object parsing.
type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social: VK: %s and FB: %s\n", u.Social.Vkontakte, u.Social.Facebook)
}

// Десериализация (последовательность байт => объект).
func main() {
	// 1. Создание файлового-дескриптора.
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}

	// cd l1_json/unmarshall && go run main.go
	defer jsonFile.Close()
	fmt.Println("File descriptor has been successfully created.")

	// 2. Десериализация json в экземпляр Go.
	var users Users

	// 2.1 Вычитываем содержимое json в виде ПОСЛЕДОВАТЕЛЬНОСТИ БАЙТ.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(byteValue)

	// 2.2 Перенос из byteValue в users == ДЕСЕРИАЛИЗАЦИЯ.
	json.Unmarshal(byteValue, &users)
	fmt.Println(users)
	for _, u := range users.Users {
		fmt.Println("====")
		PrintUser(&u)
	}
}
