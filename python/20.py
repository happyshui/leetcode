#!/usr/bin/python3

def isValid(s: str) -> bool:
    dicts = {')':'(',']':'[','}':'{'}
    stack = []
    if len(s) % 2 == 1:
        return False
    for i in s:
        if stack and i in dicts:
            if stack[-1] == dicts[i]:
                stack.pop()
            else:
                return False
        else:
            stack.append(i)
    return not stack

if __name__ == '__main__':
    str = "()"
    print(isValid(str))
