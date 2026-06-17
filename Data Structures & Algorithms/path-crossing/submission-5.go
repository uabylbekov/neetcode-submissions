func isPathCrossing(path string) bool {
    x, y := 0, 0

	seen := make(map[[2]int]struct{})
	seen[[2]int{x, y}] = struct{}{}

	for _, p := range path {
		if p == 'N' {
			y++
		} else if p == 'S' {
			y--
		} else if p == 'W' {
			x++
		} else if p == 'E' {
			x--
		}

		if _, ok := seen[[2]int{x, y}]; ok {
			return true
		}

		seen[[2]int{x,y}] = struct{}{}
	}
	return false
}