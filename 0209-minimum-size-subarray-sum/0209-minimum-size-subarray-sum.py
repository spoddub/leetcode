class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        begin = 0
        window_state = 0
        result = float('inf')

        for end in range(len(nums)):
            window_state += nums[end]

            while window_state >= target:
                window_size = end - begin + 1
                result = min(result, window_size)
                window_state -= nums[begin]
                begin += 1 #shrink
        
        if result == float("inf"):
            return 0
        
        return result