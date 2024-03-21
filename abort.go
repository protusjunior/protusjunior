package piscine

func Abort(a, b, c, d, e int) int {
	numSlice := []int{a, b, c, d, e}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				numSlice[i], numSlice[j] = numSlice[j], numSlice[i]
			}
		}
	}
	if len(numSlice)%2 == 0 {
		index := (len(numSlice) / 2) - 1
		return numSlice[index]
	} else {
		index := len(numSlice) / 2
		return numSlice[index]
	}
}
