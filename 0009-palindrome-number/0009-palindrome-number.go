func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    originalNum := x
    reversedNum := 0

    for x > 0 {
        digit := x % 10
        reversedNum = reversedNum * 10 + digit
        x = x / 10
    }

    return originalNum == reversedNum

}