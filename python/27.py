#!/usr/bin/python

def removeDuplicates(nums, val):
    if len(nums) == 0:
        return 0
    l = len(nums)
    i = 0
    for j in range(0, l):
        if nums[j] != val:
            nums[i] = nums[j]
            i += 1
    return i


if __name__ == '__main__':
    nums = [3, 2, 2, 3]
    val = 3
    print(removeDuplicates(nums, val))
