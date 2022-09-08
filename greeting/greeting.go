package main

const englisHelloPrefix = "Hello, "

func Greeting(name string) string {

	if name == "" {
		name = "World "
	}

	return englisHelloPrefix + name
}
