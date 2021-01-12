#!/usr/bin/python3

import sys

# def pa(x):
#     if x == 0:
#         return 0
#     if x == 1:
#         return 1
#     if x == 2:
#         return 2
#     else:
#         return pa(x - 1) + pa(x - 2)

def pa(x):
    if x == 0:
        return 0
    elif x == 1:
        return 1
    elif x == 2:
        return 2
    else:
        a = b = 1
        for i in range(x):
            a, b = b, a + b
        return a


if __name__ == '__main__':
    x = int(sys.argv[1])
    print(pa(x))
