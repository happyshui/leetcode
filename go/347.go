package leetcode

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	nmap := map[int]int{}
	for _, num := range nums {
		nmap[num]++
	}

	sh := &smallheap{}
	heap.Init(sh)
	for key, value := range nmap {
		heap.Push(sh, [2]int{key, value})
		if sh.Len() > k {
			heap.Pop(sh)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(sh).([2]int)[0]
	}
	return ret
}

type smallheap [][2]int

func (sh *smallheap) Push(n interface{}) {
	*sh = append(*sh, n.([2]int))
}

func (sh *smallheap) Pop() interface{} {
	before := *sh
	end := len(before)
	n := before[end-1]
	*sh = before[0 : end-1]
	return n
}

func (sh smallheap) Len() int           { return len(sh) }
func (sh smallheap) Less(i, j int) bool { return sh[i][1] < sh[j][1] }
func (sh smallheap) Swap(i, j int)      { sh[i], sh[j] = sh[j], sh[i] }
