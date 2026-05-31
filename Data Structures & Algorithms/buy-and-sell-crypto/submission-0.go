func maxProfit(prices []int) int {
    minPrice := prices[0]
    maxProfit := 0

    for _, price := range prices {
        minPrice = min(minPrice, price)
        maxProfit = max(maxProfit, price - minPrice)
    }

    return maxProfit
}
