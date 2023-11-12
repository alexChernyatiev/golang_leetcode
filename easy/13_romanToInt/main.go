package main

import "fmt"

func romanToInt(s string) int {
	res := 0
	isLeft := 0
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			nextSymbol := string(s[i+1])
			firstStatement := string(s[i]) == "I" && (nextSymbol == "V" || nextSymbol == "X")
			secondStatement := string(s[i]) == "X" && (nextSymbol == "L" || nextSymbol == "C")
			thirdStatement := string(s[i]) == "C" && (nextSymbol == "D" || nextSymbol == "M")

			if firstStatement || secondStatement || thirdStatement {
				isLeft = 1
				continue
			}
		}

		switch string(s[i]) {
		case "I":
			res += 1
		case "V":
			res += 5 - isLeft*1
		case "X":
			res += 10 - isLeft*1
		case "L":
			res += 50 - isLeft*10
		case "C":
			res += 100 - isLeft*10
		case "D":
			res += 500 - isLeft*100
		case "M":
			res += 1000 - isLeft*100
		}

		isLeft = 0
	}

	return res
}

func main() {
	fmt.Print(romanToInt("III"), "\n")
	fmt.Print(romanToInt("LVIII"), "\n")
	fmt.Print(romanToInt("MCMXCIV"), "\n")
}
