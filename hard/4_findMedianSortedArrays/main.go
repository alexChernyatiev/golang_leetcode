// You can edit this code!
// Click here and start typing.
package main

import (
	"math"
)

func printArr(nums []int) {
	for _, num := range nums {
		print(num, "->")
	}
	print("\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Merging arrays and using 2 pointers. Time: O(m+n), Space: O(m+n)
/*func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		merged = append(merged, nums1[i])
		i++
	}
	for j < len(nums2) {
		merged = append(merged, nums2[j])
		j++
	}

	if len(merged)%2 != 0 {
		return float64(merged[len(merged)/2])
	} else {
		pos := len(merged) / 2
		return float64(merged[pos]+merged[pos-1]) / 2
	}
}*/

// Using Binary Search. Time: O(log(min(m,n))), Space: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		nums2, nums1 = nums1, nums2
	}

	n, m := len(nums1), len(nums2)
	start, end := 0, n
	mergedMid := (n + m + 1) / 2

	for start <= end {
		mid := (end + start) / 2
		l1Size := mid
		l2Size := mergedMid - mid

		left1 := math.MinInt32
		if l1Size > 0 {
			left1 = nums1[l1Size-1]
		}

		left2 := math.MinInt32
		if l2Size > 0 {
			left2 = nums2[l2Size-1]
		}

		right1 := math.MaxInt32
		if l1Size < n {
			right1 = nums1[l1Size]
		}

		right2 := math.MaxInt32
		if l2Size < m {
			right2 = nums2[l2Size]
		}

		if left1 <= right2 && left2 <= right1 {
			if (m+n)%2 == 0 {
				return float64(max(left1, left2)+min(right1, right2)) / 2.0
			} else {
				return float64(max(left1, left2))
			}
		} else if left1 > right2 {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return 0
}

func main() {
	print(findMedianSortedArrays([]int{1, 3}, []int{2}))
	print(findMedianSortedArrays([]int{1, 3, 5}, []int{2, 4}))
	print(findMedianSortedArrays([]int{3}, []int{-1, -2}))
}
