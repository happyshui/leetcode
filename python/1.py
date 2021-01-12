#!/usr/bin/python3

def twoSum(nums, target):
    i = 0
    j = 0
    l = len(nums)
    for i in range(0, l):
        for j in range(1, l):
            if nums[i] + nums[j] == target and i != j:
                return (i, j)


if __name__ == '__main__':
    nums = [2, 5, 5, 11]
    target = 10
    print(twoSum(nums, target))

