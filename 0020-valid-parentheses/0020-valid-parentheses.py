class Solution:
    def isValid(self, s: str) -> bool:
        pair = {
            "(": ")",
            "{": "}",
            "[": "]",
        }
        stack = []
        
        for c in s:
            if c in pair:
                stack.append(c)
            else:
                if not stack:
                    return False
                
                prev = stack.pop()
                if c != pair[prev]:
                    return False
        
        return len(stack) == 0