func majorityElement(nums []int) int {
    countNums := make(map[int]int)

	for _, num := range nums {
		countNums[num]++
	}

	for key, val := range countNums {
		if val > len(nums) / 2 {
			return key
		}
	}

	return -1
}
