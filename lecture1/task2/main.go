package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	str := strs[0]
	count := 0

	for i := range str {
		for j := range strs {
			if str[i] != strs[j][i] {
				return str[0:count]
			}
		}
		count += 1
	}

	return str[0:count]
}

func main() {
	strs := []string{"flower", "flo", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
