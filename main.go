package main

import "fmt"

func main() {

	input := []int{1, 3, 5, 6}
	target := 5
	result := searchInsert(input, target)
	fmt.Println(result)

}

func searchInsert(nums []int, target int) int {
	var result int
	last := len(nums) - 1
	for i := 0; i < len(nums); i++ {

		if target > nums[last] {
			fmt.Println(last + 1)
			result = last + 1
			break
		}

		if nums[i] == target {
			fmt.Println(i)
			result = i
			break
		} else if nums[i] > target {
			fmt.Println(i)
			result = i
			break
		} else {

		}
	}
	return result
}
