func canConstruct(ransomNote string, magazine string) bool {
	seen := make(map[rune]int)

	for _, char := range magazine {
		seen[char]++
	}

	for _, char := range ransomNote {
		if seen[char] == 0 {
			return false
		}

		seen[char]--
	}

	return true
}