func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		if target == numbers[l] + numbers[r] {
			return []int { l + 1, r + 1}
		} else if target < numbers[l] + numbers[r] {
			r--
		} else {
			l++
		}
	}

	return []int{-1, -1}
}
