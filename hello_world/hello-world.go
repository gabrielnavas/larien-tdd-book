package helloworld

const prefixHelloPortuguese = "Olá, "

func HelloWorld(name string) string {
	if name == "" {
		return "Hello World"
	}
	return prefixHelloPortuguese + name
}
