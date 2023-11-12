package main

func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		//print(i, " ", j, " ", nums[i], " ", nums[j], "\n")
		if i != j && nums[i] != nums[j] {
			j++
			nums[j], nums[i] = nums[i], nums[j]

		}
	}

	//printArr(nums)
	return j + 1
}

func printArr(digits []int) {
	for i := 0; i < len(digits); i++ {
		println(digits[i])
	}
}

func main() {
	print("\n", removeDuplicates([]int{1, 1, 2, 2, 2, 3, 4, 5, 5}))
	//print("\n", removeDuplicates([]int{1, 1, 2}))
}
