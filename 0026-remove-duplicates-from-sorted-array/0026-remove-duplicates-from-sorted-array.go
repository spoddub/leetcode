func removeDuplicates(nums []int) int {
	k := 1

	for i := range nums {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
