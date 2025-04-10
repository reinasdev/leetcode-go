package easy

// Big O: O(n)
func isValid(s string) bool {
	var stack []uint8
	for i := range s {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}

			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// Big O: O(n)
func isValidOptimized(s string) bool {
	var (
		stack   []uint8
		closing = map[uint8]uint8{')': '(', '}': '{', ']': '['}
	)

	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 || stack[len(stack)-1] != closing[s[i]] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
