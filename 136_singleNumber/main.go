package main

//Solution with XOR
//'x^y=y^x', 'x^0=x' and 'x^x=0' so we can remove pairs of numbers and receive result
func singleNumber(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum ^ nums[i]
	}

	return sum
}

//Solution with using Map time: O(n), memory: O(n)
/*func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			delete(m, nums[i])
		} else {
			m[nums[i]] = 1
		}
	}

	for key := range m {
		return key
	}
}*/

func main() {
	print("\n", singleNumber([]int{1, 2, 2, 1, 4, 3, 4, 5, 5}))
}
