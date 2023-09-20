package main

import "fmt"

func main() {
	nums := []int{2, 7, 8, 10}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)

    for i := 0; i < len(nums); i ++ {
        val, ok := mp[target-nums[i]]
        if ok {
            return []int{i, val}
        }else{
            mp[nums[i]] = i
        }
    }

    return []int{-1, -1}
}