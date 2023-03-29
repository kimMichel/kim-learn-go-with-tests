package main

import "fmt"

const prefixHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefixHello + name
}

func main() {
	fmt.Println(Hello("Kim"))
}
