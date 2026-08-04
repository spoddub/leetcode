func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, word := range strs[1:] {
		i := 0

		for i < len(prefix) &&
			i < len(word) &&
			prefix[i] == word[i] {
			i++
		}

		prefix = prefix[:i]

		if prefix == "" {
			return ""
		}
	}

	return prefix
}