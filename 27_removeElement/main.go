package main

//Hard and very fast alg
/*func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			count++
			continue
		}

		if i >= len(nums)-1 {
			continue
		}

		j := i + 1
		found := false
		for j <= len(nums)-1 {
			if nums[j] != val {
				found = true
				break
			}
			j++
		}

		print(i, j, found, nums[i], nums[j], j < len(nums)-1, "\n")
		if !found {
			continue
		}

		buf := nums[i]
		nums[i] = nums[j]
		nums[j] = buf
		count = i + 1
	}

	printArr(nums)
	return count
}*/

//Simple and more slowly alg
func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	printArr(nums)
	return j
}

func printArr(digits []int) {
	for i := 0; i < len(digits); i++ {
		println(digits[i])
	}
}

func main() {
	print("\n", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	print("\n", removeElement([]int{4, 5}, 4))
}
