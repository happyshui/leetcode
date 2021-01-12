#!/usr/bin/python3

def maxArea(lists):
    value, start, end = 0, 0, len(lists)-1
    while start < end:
        width = end - start
        high = 0
        if lists[start] < lists[end]:
            high = lists[start]
            start += 1
        else:
            high = lists[end]
            end -= 1

        temp = width * high
        print(width, high)
        print(start, end)
        print(temp, "tt")
        if temp > value:
            value = temp
    
    return value


if __name__ == '__main__':
    lists = [9, 8, 6, 2, 5, 4, 8, 3, 7, 8]
    print(maxArea(lists))
