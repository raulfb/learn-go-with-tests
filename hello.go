package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const galicianHelloPrefix = "Ola, "
const spanish = "Spanish"
const french = "French"
const galician = "Galician"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case galician:
		prefix = galicianHelloPrefix
	}

	return prefix + name
}
func main() {
	fmt.Println(Hello("world", ""))
}
