#!/usr/bin/python3

def merge(nums1, m, nums2, n):
    p1 = m - 1
    p2 = n - 1

    p = m + n -1
    
    while p1 >= 0 and p2 >= 0:
        if nums1[p1] < nums2[p2]:
            nums1[p] = nums2[p2]
            p2 -= 1
        else:
            nums1[p] = nums1[p1]
            p1 -= 1
        p -= 1

    nums1[:p2 + 1] = nums2[:p2 + 1]
    return nums1


def merge1(nums1, m, nums2, n):
    nums1[:] = sorted(nums1[:m] + nums2)
    return nums1



if __name__ == '__main__':
    nums1 = [1,2,3,0,0,0]
    m = 3
    nums2 = [2, 5, 6]
    n = 3
    print(merge(nums1, m, nums2, n))

    nums3 = [1,2,3,0,0,0]
    m = 3
    nums4 = [2, 5, 6]
    n = 3
    print(merge1(nums3, m, nums4, n))

