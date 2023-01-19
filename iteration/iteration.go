package iteration

func Repeat(character string, iterationMany int) string {
	characters := ""
	for i := 0; i < iterationMany; i++ {
		characters += character
	}
	return characters
}
