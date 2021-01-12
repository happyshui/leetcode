#!/usr/bin/python3

import sys


def lengthOfLastWorld(s):
    tmp = s.split(' ')
    num = 0
    for _ in tmp[::-1]:
        if _ != '':
            num += 1
            return len(_)
        else:
            num += 0
    if num == 0:
        return 0



if __name__ == '__main__':
    s = sys.argv[1]
    print(lengthOfLastWorld(s))

