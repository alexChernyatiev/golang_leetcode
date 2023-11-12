package main

//time: O(2n); space: O(n) for call stack
/*func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}*/

//O(n) time and O(1) space
func climbStairs(n int) int {
	res := 0

	next := 0
	secondNext := 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			res = 1
		} else if i == 2 {
			res = 2
		} else {
			res = next + secondNext
		}

		next = secondNext
		secondNext = res
	}

	return res
}

func main() {
	print(climbStairs(3))  //3
	print(climbStairs(4))  //5
	print(climbStairs(5))  //8
	print(climbStairs(45)) //1836311903
}
