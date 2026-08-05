func firstUniqChar(s string) int {
	counter := make(map[rune]int)

	for _, char := range s {
		counter[char]++
	}

	for i, char := range s {
		if counter[char] == 1 {
			return i
		}
	}

	return -1
}