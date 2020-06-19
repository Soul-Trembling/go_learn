package main

import "fmt"

func twoSum1(nums []int, target int) []int {
	var result []int
	if len(nums) < 2 {
		return result
	}
	var exist map[int]int
	exist = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		exist[nums[i]] = i
	}

	for num, index := range exist {
		temp := target - num
		tempIndex := exist[temp]
		if tempIndex > 0 {
			result = append(result, index)
			result = append(result, tempIndex)
		}
	}
	return result
}

func twoSum(nums []int, target int) []int {
	var result []int
	if len(nums) < 2 {
		return result
	}
	var exist map[int]int
	exist = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if index, ok := exist[target-nums[i]]; ok {
			result = append(result, index)
			result = append(result, i)
		}
		exist[nums[i]] = i
	}

	return result
}

func main() {
	nums := []int{2, 11, 7, 15}
	var target = 9
	fmt.Printf("%v", twoSum(nums, target))
	println()
	fmt.Printf("%v", twoSum1(nums, target))
}
