package general

import (
	"math/bits"
	"math/rand"
)

func Choice(nodes []int, k int) []int {
	res := make([]int, 0, k)

	var val int
	for i := 0; i < k; i++ {
		val = rand.Intn(len(nodes))
		res = append(res, nodes[val-1])
	}
	return res
}

func BitLength(num int) int {
	return bits.Len(uint(num))
}

func Contains[T comparable](elems []T, value T) bool {
	for _, item := range elems {
		if item == value {
			return true
		}
	}
	return false
}

func BinarySearchClosest(arr []int, target int, n int) []int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		curNodeId := arr[mid]
		if curNodeId > target-n && curNodeId < target+n {
			return findClosest(arr, target, mid, n)
		} else if curNodeId < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return findClosest(arr, target, left, n)
}

func findClosest(arr []int, target int, index int, n int) []int {
	result := make([]int, 0, n)
	left, right := index-1, index+1
	for left >= 0 && right < len(arr) && len(result) < n {
		if target-arr[left] < arr[right]-target {
			result = append(result, arr[left])
			left--
		} else {
			result = append(result, arr[right])
			right++
		}
	}
	for left >= 0 && len(result) < n {
		result = append(result, arr[left])
		left--
	}
	for right < len(arr) && len(result) < n {
		result = append(result, arr[right])
		right++
	}
	return result
}
