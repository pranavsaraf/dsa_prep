/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].
You must write an algorithm with O(log n) runtime complexity.
Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:
Input: nums = [], target = 0
Output: [-1,-1]
*/
package main

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	start := binarySearch(nums, target, true)
	end := binarySearch(nums, target, false)
	ans[0] = start
	ans[1] = end
	return ans
}

func binarySearch(nums []int, target int, isStartIndex bool) int {
	ans := -1
	n := len(nums)
	start := 0
	end := n - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			ans = mid
			if isStartIndex {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return ans
}
