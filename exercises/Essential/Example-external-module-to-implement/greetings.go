package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nombre vacio")
	}
	message := fmt.Sprintf(RandomFormat(), name)
	return message, nil
}

func RandomFormat() string {
	formats := []string{
		"Hola, %v. Bienvenido",
		"Que gusto verte, %v",
		"Saludos %v",
	}

	return formats[rand.Intn(len(formats))]
}

func SayHello(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message + "\n"
	}

	return messages, nil
}
