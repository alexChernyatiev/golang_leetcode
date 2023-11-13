package main

func printArr(nums1 []int) {
	for _, num := range nums1 {
		print(num, "->")
	}
	print("\n")
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j := m-1, n-1
	for k := m + n - 1; k >= 0; k-- {
		if j < 0 {
			break
		}
		if i < 0 || nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j -= 1
		} else {
			nums1[k], nums1[i] = nums1[i], nums1[k]
			i -= 1
		}
	}
	return nums1
}

func main() {
	printArr(merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3))
	printArr(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	printArr(merge([]int{2, 4, 6, 0, 0, 0}, 3, []int{1, 3, 7}, 3))
	printArr(merge([]int{1}, 1, []int{0}, 0))
	printArr(merge([]int{0}, 0, []int{1}, 1))
}
