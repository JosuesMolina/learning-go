package main

import (
	"fmt"
	"log"

	"github.com/JosuesMolina/greetings"
)

func main() {
	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Josue")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Josue", "Nubia", "Apache", "Oddie"}
	messages, err := greetings.SayHello(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

// Se utilizan los siguientes comandos para poder hacer la referencia del package local:
// go mod edit -replace github.com/JosuesMolina/greetings=../Example-external-module-to-implement
// go mod tidy
