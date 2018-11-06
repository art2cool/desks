package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	// Enter solution here
	// := 1
	arr := strings.Split(str, "")
	for _, l := range arr {
		if l == "a" || l == "e" || l == "i" || l == "o" || l == "u" {
			count = count + 1
		}
	}

	return count
}

func main() {
	fmt.Println(GetCount("vova"))
}
