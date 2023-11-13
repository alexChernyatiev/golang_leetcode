package main

func strStr(haystack string, needle string) int {
	i, j, k := 0, 0, -1
	for i < len(needle) && j < len(haystack) {
		if needle[i] == haystack[j] {
			if k == -1 {
				k = j
			}
			i, j = i+1, j+1
		} else {
			if k != -1 {
				j, k = k+1, -1
			} else {
				j += 1
			}

			i = 0
		}

	}

	if i == len(needle) {
		return j - i
	} else {
		return -1
	}
}

func main() {
	print(strStr("sadbutsad", "sad"))
	print(strStr("leetcode", "leeto"))
	print(strStr("mississippi", "issip"))
	print(strStr("aabaabbbaabbbbabaaab", "abaa"))
}
