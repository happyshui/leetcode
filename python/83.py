#!/usr/bin/python3

def deleteDuplicates(head):
    if head is None or head.next is None: return head
    cur = head
    while cur != None and  cur.next != None:
        if cur.val == cur.next.val:
            cur.next = cur.next.next
        else:
            cur = cur.next
    return head


if __name__ == '__main__':
    head = [1, 1, 2]
    print(deleteDuplicates(head))
