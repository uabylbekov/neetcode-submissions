func maxArea(heights []int) int {
    l := 0
    r := len(heights) - 1
    maxAmount := 0

    for l < r {
        minHeight := min(heights[l], heights[r])
        maxAmount = max(maxAmount, minHeight * (r - l))

        if heights[l] > heights[r] {
            r--
        } else {
            l++
        }
    }

    return maxAmount
}
