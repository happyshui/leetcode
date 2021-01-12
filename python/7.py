#!/usr/bin/python3


import sys

def reverse(x):
    tmp = 0
    while x != 0:
        if x > 0:
            tmp = tmp * 10 + x % 10
            x = x // 10
        if x < 0:
            tmp = tmp * 10 + x % -10
            x = -(x // -10)
    if tmp < -2**31 or tmp > 2**31-1:
        return 0
    return tmp


if __name__ == '__main__':
    num = int(sys.argv[1])
    print(reverse(num))

