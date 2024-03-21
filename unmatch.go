package piscine

func Unmatch(nums []int) int {
	for _, nb := range nums {
		count := 0
		for _, num := range nums {
			if nb == num {
				count++
			}
		}
		if count%2 != 0 {
			return nb
		}

	}
	return -1
}
