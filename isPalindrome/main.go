package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var s []int
	num := x
	for num > 0 {
		s = append(s, num%10)
		num = num / 10
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(123))
}
