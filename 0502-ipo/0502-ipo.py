class Solution:
    def findMaximizedCapital(self, k: int, w: int, profits: List[int], capital: List[int]) -> int:
        projects = sorted(list(zip(capital, profits)))
        can_afford = []
        checked = 0

        for _ in range(k):
            while checked < len(projects) and projects[checked][0] <= w:
                heapq.heappush(can_afford, -projects[checked][1])
                checked += 1

            if not can_afford:
                break
            
            w += -heapq.heappop(can_afford)

        return w