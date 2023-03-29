package main

import "fmt"

func Hello() string {
	return "Hello World!"
}

func HelloWithName(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(Hello())
}
