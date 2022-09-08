package main

const polish = "Polish"
const norwegian = "Norwegian"
const englishHelloPrefix = "Hello, "
const polishHelloPrefix = "Cześć, "
const norwegianHelloPrefix = "Hallo, "

func Greeting(name, lang string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) string {
	switch lang {
	case polish:
		return polishHelloPrefix
	case norwegian:
		return norwegianHelloPrefix
	default:
		return englishHelloPrefix
	}
}
