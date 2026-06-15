func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([][]int, 0)


	for i, temperature := range temperatures {	
		for len(stack) > 0 && temperature > stack[len(stack) - 1][0] {
			result[stack[len(stack)-1][1]] = i - stack[len(stack) - 1][1]
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, []int{temperature, i})
	}

	return result
}
