class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        begin = 0
        window_state = 0
        result = 0

        for end in range(len(nums)):
            if nums[end] == 0:
                window_state += 1
            while  window_state > k:
                
                if nums[begin] == 0:
                    window_state -= 1
                begin += 1  # shrink
            result = max(result, end - begin + 1)
        return result