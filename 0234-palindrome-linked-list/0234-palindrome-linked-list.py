# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        def middle(head):
            slow = head
            fast = head

            while fast and fast.next:
                fast = fast.next.next
                slow = slow.next

            return slow

            
        def reverse(head):
            prev = None
            current = head
            while current:
                tmp = current.next
                current.next = prev
                prev = current
                current = tmp
            
            return prev

        mid = middle(head)
        second = reverse(mid)
        first = head
        
        while first and second:
            if first.val != second.val:
                return False
            first = first.next
            second = second.next
        return True