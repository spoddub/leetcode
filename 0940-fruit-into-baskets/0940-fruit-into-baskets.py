class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        begin = 0
        window_state = defaultdict(int)
        result = 0

        for end in range(len(fruits)):
            window_state[fruits[end]] += 1

            while len(window_state) > 2:
                window_state[fruits[begin]] -= 1
                if window_state[fruits[begin]] == 0:
                    del window_state[fruits[begin]]
                begin += 1  # shrink
            result = max(result, end - begin + 1)
        return result