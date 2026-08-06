func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]bool)

	nums3 := []int{}

	for _, n := range nums1 {
		seen[n] = true
	}

	for _, k := range nums2 {
		if _, exists := seen[k]; exists {
			nums3 = append(nums3, k)
			delete(seen, k)
		}
	}

	return nums3
}