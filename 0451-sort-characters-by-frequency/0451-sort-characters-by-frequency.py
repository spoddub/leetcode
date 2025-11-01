class Solution:
    def frequencySort(self, s: str) -> str:
        c = Counter(s)
        max_heap = []

        for ch, value in c.items():
            heapq.heappush(max_heap, (-value, ch*value))
        
        result = ""

        while max_heap:
            result += heapq.heappop(max_heap)[1]
        
        return result
        