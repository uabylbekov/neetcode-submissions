import "maps"

func maxDifference(s string) int {
	countS := make(map[rune]int)

	for _, char := range s {
		countS[char]++
	}

	maxOdd := math.MinInt
	minEven := math.MaxInt

	for val := range maps.Values(countS) {
		if val % 2 != 0 {
			maxOdd = max(maxOdd, val)
		} else {
			minEven = min(minEven, val)
		}
	}


	return maxOdd - minEven
}
