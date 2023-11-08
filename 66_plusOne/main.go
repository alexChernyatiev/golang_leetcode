package main

func plusOne(digits []int) []int {
	k := true
	for i := len(digits) - 1; i >= 0; i-- {
		var plus int
		if k == true {
			plus = 1
		} else {
			plus = 0
		}
		sum := digits[i] + plus
		digits[i] = sum % 10
		k = sum >= 10
	}

	if k == true {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func printArr(digits []int) {
	for i := 0; i < len(digits); i++ {
		println(digits[i])
	}
}

func main() {
	printArr(plusOne([]int{9, 9, 9}))
}
