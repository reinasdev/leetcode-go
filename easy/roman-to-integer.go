package easy

// Big O: O(n)
func romanToInt(s string) int {
	var (
		vals   = map[uint8]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
		result int
	)

	for j := range s {
		result += vals[s[j]]

		if j != 0 {
			if vals[s[j]] > vals[s[j-1]] {
				result -= 2 * vals[s[j-1]]
			}
		}
	}

	return result
}
