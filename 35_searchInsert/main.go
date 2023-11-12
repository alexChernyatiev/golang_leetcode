package main

func searchInsert(nums []int, target int) int {
	arr := nums

	plus := 0
	for true {
		mid := len(arr) / 2
		if arr[mid] == target {
			return mid + plus
		}

		if len(arr) == 1 {
			if arr[0] >= target {
				return plus
			} else {
				return plus + 1
			}
		}

		if arr[mid] >= target {
			arr = arr[:mid]
		} else {
			plus += mid
			arr = arr[mid:]
		}
	}

	return -1
}

func main() {
	print(searchInsert([]int{1}, 1))          //0
	print(searchInsert([]int{1}, 0))          //0
	print(searchInsert([]int{1}, 2))          //1
	print(searchInsert([]int{1, 3, 5}, 5))    //2
	print(searchInsert([]int{1, 3, 5, 6}, 0)) //0
	print(searchInsert([]int{1, 3, 5, 6}, 2)) //1
	print(searchInsert([]int{1, 3, 5, 6}, 7)) //4
}
