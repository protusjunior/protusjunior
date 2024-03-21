package piscine

func ToUpper(s string) string {
	ToUpperRune := []rune(s)
	NewString := ""
	for i := 0; i < len(s); i++ {
		a := ToUpperRune[i]
		if a >= 97 && a < 123 {
			a -= 32
		}
		NewString += string(a)
	}
	return NewString
}
