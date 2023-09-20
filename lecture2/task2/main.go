package main

import "fmt"

func main() {
	fmt.Print(intToRoman(58))

}
func intToRoman(num int) string {
	var res string

	sl := []struct {
		Key int
		Val string
	}{
		{
			Key: 1000,
			Val: "M",
		},
		{
			Key: 900,
			Val: "CM",
		},
		{
			Key: 500,
			Val: "D",
		},
		{
			Key: 400,
			Val: "CD",
		},
		{
			Key: 100,
			Val: "C",
		},
		{
			Key: 90,
			Val: "XC",
		},
		{
			Key: 50,
			Val: "L",
		},
		{
			Key: 40,
			Val: "XL",
		},
		{
			Key: 10,
			Val: "X",
		},
		{
			Key: 9,
			Val: "IX",
		},
		{
			Key: 5,
			Val: "V",
		},
		{
			Key: 4,
			Val: "IV",
		},
		{
			Key: 1,
			Val: "I",
		},
	}

	for _, kvp := range sl {
		n := num / kvp.Key
		for i := 0; i < n; i++ {
			res += kvp.Val
		}
		num %= kvp.Key
	}

	return res
}
