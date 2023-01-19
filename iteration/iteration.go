package iteration

const iterationMany = 5

func Repeat(character string) string {
	characters := ""
	for i := 0; i < iterationMany; i++ {
		characters += character
	}
	return characters
}
