package iteration

func Repeat(character string) string {
	characters := ""
	for i := 0; i < 5; i++ {
		characters += character
	}
	return characters
}
