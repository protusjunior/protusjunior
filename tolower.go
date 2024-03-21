package piscine

func ToLower(s string) string {
	ToLowerRune := []rune(s)
	NewString := ""
	for i := 0; i < len(s); i++ {
		a := ToLowerRune[i]
		if a >= 65 && a < 91 {
			a += 32
		}
		NewString += string(a)
	}
	return NewString
}
