package main

import (
	"fmt"
)

const (
    russian = "Russian"
    spanish = "Spanish"
    french = "French"
    russianHelloPrefix = "Привет, "
    frenchHelloPrefix = "Bonjour, "
    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case spanish:
        prefix = spanishHelloPrefix
    case french:
        prefix = frenchHelloPrefix
    case russian:
        prefix = russianHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}

func main() {
    hello := Hello("Chris chan", "English")
	fmt.Println(hello)
}
