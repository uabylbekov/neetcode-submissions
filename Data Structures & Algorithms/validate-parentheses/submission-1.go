func isValid(s string) bool {
    openersToClosers := map[rune]rune {
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := make([]rune, 0)

	for _, char := range s {
		if _, ok := openersToClosers[char]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if openersToClosers[last] != char {
				return false
			}
		}
	}

	return len(stack) == 0
}
