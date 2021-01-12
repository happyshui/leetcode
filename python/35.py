#!/usr/bin/python3

def searchInsert(nums, target):
    start = 0
    if target > nums[len(nums)-1]:
        return len(nums)
    for i in range(len(nums)):
        if target == nums[i]:
            return i
        elif target > nums[i] and target < nums[i+1]:
            start = i + 1
    return start


if __name__ == '__main__':
    nums = [1,3,5,6]
    target = 5
    print(searchInsert(nums, target))
