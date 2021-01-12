#!/usr/bin/python3

import sys


def palindrome(num):
    if num < 0:
        return False
    if str(num) == str(num)[::-1]:
        return True
    else:
        return False


if __name__ == '__main__':
    num = int(sys.argv[1])
    print(palindrome(num))

