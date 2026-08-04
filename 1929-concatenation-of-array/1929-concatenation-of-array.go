func getConcatenation(nums []int) []int {
    n := len(nums)
    n2 := n*2
    newNums := make([]int, n2, n2)

    for i := range nums {
        newNums[i] = nums[i]
        newNums[i+n] = nums[i]
    }

    return newNums
}
