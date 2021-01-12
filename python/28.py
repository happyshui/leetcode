#!/usr/bin/python3

def strStr(haystack, needle):
    ln, lh = len(needle), len(haystack)
    for i in range(lh - ln + 1):
        if haystack[i:i+ln] == needle:
            return i
    return -1


if __name__ == '__main__':
    haystack = ""
    needle = ""
    print(strStr(haystack, needle))
