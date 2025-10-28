# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def largestValues(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        
        result = []
        q = deque([root])

        while q:
            level_size = len(q)
            level_max = float('-inf')

            for _ in range(level_size):
                node = q.popleft()
                level_max = max(level_max, node.val)

                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
            
            result.append(level_max)
        
        return result