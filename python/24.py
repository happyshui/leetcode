#!/usr/bin/python3

def swapPairs(self, head: ListNode) -> ListNode:
   cur, cur.next = self, head
   while cur.next and cur.next.next:
       a = cur.next
       b = a.next
       cur.next, b.next, a.next = b, a, b.next
       cur = a
   return self.next
