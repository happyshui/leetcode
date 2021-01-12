#!/usr/bin/python

def removeDuplicates(nums):
    if len(nums) == 0:
        return 0
    l = len(nums)
    i = 0
    for j in range(1, l):
        if nums[j] != nums[i]:
            i += 1
            nums[i] = nums[j]
    print(nums)
    return i + 1


if __name__ == '__main__':
    nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
    print(removeDuplicates(nums))
