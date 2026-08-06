func isValid(s string) bool {
	p := []rune{}

	for _, j := range s {
		switch j {
		case '[', '(', '{':
			p = append(p, j)

		case ']':
			if len(p) == 0 || p[len(p)-1] != '[' {
				return false
			}
			p = p[:len(p)-1]

		case ')':
			if len(p) == 0 || p[len(p)-1] != '(' {
				return false
			}
			p = p[:len(p)-1]
		case '}':
			if len(p) == 0 || p[len(p)-1] != '{' {
				return false
			}
			p = p[:len(p)-1]
		}
	}

	return len(p) == 0
}