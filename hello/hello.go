package main

import (
	"fmt"
)

const prefixHello = "Hello, "
const prefixOla = "Olá, "
const prefixHola = "Hola, "
const prefixBonjour = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixGreeting(language) + name
}

func prefixGreeting(language string) (prefix string) {
	switch language {
	case "portuguese":
		prefix = prefixOla
	case "spanish":
		prefix = prefixHola
	case "french":
		prefix = prefixBonjour
	default:
		prefix = prefixHello
	}
	return
}

func main() {
	fmt.Println(Hello("Kim", ""))
}
