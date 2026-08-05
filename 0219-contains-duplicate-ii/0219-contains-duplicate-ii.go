func containsNearbyDuplicate(nums []int, k int) bool {
	lastIdx := make(map[int]int)

	for i, num := range nums {
		if prevIdx, exists := lastIdx[num]; exists {
			if i-prevIdx <= k {
				return true
			}
		}

		lastIdx[num] = i
	}

	return false
}