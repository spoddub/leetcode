class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        result = set()
        n = len(nums)
        for i in range(n):
            target = -nums[i]
            left = i + 1
            right = n - 1
            while left < right:
                current_sum = nums[left] + nums[right]
                if current_sum == target:
                    result.add((nums[i], nums[left], nums[right]))
                if current_sum > target:
                    right -= 1
                else:
                    left += 1
        return list(result)