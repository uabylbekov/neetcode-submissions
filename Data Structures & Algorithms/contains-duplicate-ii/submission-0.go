func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, num := range nums {
		if val, ok := seen[num]; ok {
			if i - val <= k {
				return true
			}
		}
		seen[num] = i
	}

	return false
}
