func longestConsecutive(nums []int) int {
	seen := make(map[int]struct{})

	for _, num := range nums {
		seen[num] = struct{}{}
	}

	longest := 0 

	for _, num := range nums {
		if _, ok := seen[num - 1]; ok {
			continue
		} 
		length := 1
		for {
			if _, exists := seen[num + length]; exists {
				length++
			} else {
				break
			}
		}
		longest = max(longest, length)
	}

	return longest
}
