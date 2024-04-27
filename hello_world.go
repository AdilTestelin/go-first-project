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

	switch language {
	case "Spanish":
		return helloPrefixInSpanish + name + " !"
	case "French":
		return helloPrefixInFrench + name + " !"
	default:
		return helloPrefixInEnglish + name + " !"
	}
}

func main() {
	fmt.Println(Hello("Adil", "English"))
}
