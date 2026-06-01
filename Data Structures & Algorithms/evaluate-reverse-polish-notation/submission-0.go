func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		if token == "+" {
			first := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack = append(stack, first + second)
		} else if token == "-" {
			first := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack = append(stack, second - first)
		} else if token == "*" {
			first := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack = append(stack, first * second)
		} else if token == "/" {
			first := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack = append(stack, second / first)
		} else {
			converted, _ := strconv.Atoi(token)
			stack = append(stack, converted)
		}
	}

	return stack[0]
}
