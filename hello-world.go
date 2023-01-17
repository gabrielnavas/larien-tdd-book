package main

const prefixHelloPortuguese = "Olá, "

func HelloWorld(name string) string {
	if name == "" {
		return "Hello World"
	}
	return prefixHelloPortuguese + name
}

func main() {
	println(HelloWorld("John"))
}
