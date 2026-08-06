func maxProfit(prices []int) int {
	minPrice := prices[0]
	maximumProfit := 0
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		}

		if maximumProfit < price-minPrice {
			maximumProfit = price - minPrice
		}
	}

	return maximumProfit
}