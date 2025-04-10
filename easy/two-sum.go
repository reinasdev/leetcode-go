package easy

import "sort"

// Big O: O(n^2)
func twoSumBruteForce(nums []int, target int) []int {
	for i, v := range nums {
		for j, k := range nums[:i] {
			if v+k == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// Big O: O(n + log n)
func twoSumSorting(nums []int, target int) []int {
	type pair struct {
		value int
		index int
	}

	pairs := make([]pair, len(nums))
	for i := 0; i < len(nums); i++ {
		pairs[i] = pair{value: nums[i], index: i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	i := 0
	f := len(nums) - 1
	for i < f {
		switch {
		case pairs[i].value+pairs[f].value == target:
			return []int{pairs[i].index, pairs[f].index}
		case pairs[i].value+pairs[f].value > target:
			f--
		case pairs[i].value+pairs[f].value < target:
			i++
		}
	}

	return nil
}

// Big O: O(n)
func twoSumHashmap(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, v := range nums {
		diff := target - v

		if p, ok := hashmap[diff]; ok {
			return []int{p, i}
		}

		hashmap[v] = i
	}

	return nil
}
