func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		needed := target - num

		if j, exists := seen[needed]; exists {
			return []int{j, i}
		}

		seen[num] = i
	}

	return []int{}
}