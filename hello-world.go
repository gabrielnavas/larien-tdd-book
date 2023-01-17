package main

const prefixHelloPortuguese = "Olá, "

func HelloWorld(name string) string {
	return prefixHelloPortuguese + name
}

func main() {
	println(HelloWorld("John"))
}
