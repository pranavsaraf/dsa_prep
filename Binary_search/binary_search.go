// Normal Binary Search
// https://leetcode.com/problems/binary-search/description/

// Time Complexity: O(log n)
// Space Complexity: O(1)
package main

func search(nums []int, target int) int {
	n := len(nums)
	start := 0
	end := n - 1
	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
