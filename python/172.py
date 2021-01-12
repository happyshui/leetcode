#!/usr/bin/python3

def trailingZeroes(n: int) -> int:
    s = 0
    num = 0
    while n > 0 :
        print(n)
        n = n // 5
        num += n
    return num



if __name__ == '__main__':
    num = 126
    print(trailingZeroes(num))
