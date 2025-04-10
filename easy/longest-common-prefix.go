package easy

import "sort"

// Big O: O(n * m)
func longestCommonPrefix(strs []string) string {
	var result []uint8 = []uint8(strs[0])

	for _, v := range strs[1:] {
		if len(result) > len(v) {
			result = result[:len(v)]
		}

		for j := range v {
			if len(result) <= j {
				continue
			}

			if v[j] != result[j] {
				result = result[:j]
			}
		}
	}

	return string(result)
}

// Big O: O(n ∗ m ∗ log n)
func longestCommonPrefixUnoptimized(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	for i := range strs[0] {
		if strs[0][i] != strs[len(strs)-1][i] {
			return strs[0][:i]
		}
	}

	return strs[0]
}
