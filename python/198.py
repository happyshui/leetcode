#!/usr/bin/python3

def rob(nums):
    homes = len(nums)
    money0 = nums[0]
    i = 0
    while i < homes-3:
        print(i)
        if nums[i] + nums[i + 2] < nums[i] + nums[i + 3]:
            money0 += nums[i+3]
            print(money0)
            i += 3
            if i >= len(nums) - 3:
                money0 += nums[i+2]
                break
        else:
            money0 += nums[i+2]
            print(money0, '+2')
            i += 2
            print(i, '11')
            if i >= len(nums)-3:
                money0 += nums[i+2]
                print(money0, '++2')
                break
    money1 = nums[1]
    i = 1
    while i < homes-3:
        if nums[i] + nums[i + 2] < nums[i] + nums[i + 3]:
            money1 += nums[i+3]
            i += 3
            if i >= len(nums):
                money1 += nums[i+2]
                break
        else:
            money1 += nums[i+2]
            i += 2    
            if i >= len(nums):
                money1 += nums[i+2]
                break
    print(money0, money1)
    return max(money0, money1)


if __name__ == '__main__':
    nums = [2,7,9,3,1]
    # nums = [2,7,9,6,1]
    # nums = [2,1,1,2]
    print(rob(nums))
