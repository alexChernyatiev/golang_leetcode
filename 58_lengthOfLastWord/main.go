package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	str := strings.Fields(s)
	return len(str[len(str)-1])
}

func main() {
	fmt.Println(lengthOfLastWord(" Hello    world!  "))
}
