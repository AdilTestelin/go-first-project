package main

import "fmt"

const helloPrefixInEnglish = "Hello "
const helloPrefixInSpanish = "Hola "
const helloPrefixInFrench = "Bonjour "
const defaultName = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingsPrefix(language) + name + " !"
}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = helloPrefixInSpanish
	case "French":
		prefix = helloPrefixInFrench
	default:
		prefix = helloPrefixInEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("Adil", "English"))
}
