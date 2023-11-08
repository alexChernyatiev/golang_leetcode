package main

import (
	"strings"
)

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	} else {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}

	res := ""
	k := false
	for i := len(a) - 1; i >= 0; i-- {
		summ := "0"
		newK := false
		if (a[i] == '1' && b[i] == '0') || (a[i] == '0' && b[i] == '1') {
			summ = "1"
			if k == true {
				newK = true
			}
		}

		if a[i] == '1' && b[i] == '1' {
			newK = true
		}

		if k == true {
			if summ == "1" {
				summ = "0"
			} else {
				summ = "1"
			}
		}

		res = summ + res
		k = newK
	}

	if k == true {
		res = "1" + res
	}

	return res
}

func main() {
	print("\n", addBinary("11", "1"))
	print("\n", addBinary("1010", "1011"))
}
