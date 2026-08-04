func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[byte]int)

	for i := range s {
		count[s[i]]++
	}

	for i := range t {
		count[t[i]]--
		
		if count[t[i]] < 0 {
			return false
		}
	}

	return true

}
