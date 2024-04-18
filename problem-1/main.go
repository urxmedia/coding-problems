// len = 3
// 0 | 1 | 2 => 1, 2, 1
// 0 + 3 | 1 + 3 | 2 + 3 => 1, 2, 1
// 0, 1, 2, 3, 4, 5 -- n = 3, 2 * n = 6

package main

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	if n >= 1 && n <= 1000 {
		for i := range ans {
			if i >= n {
				ans[i] = ans[i-n]
			} else {
				ans[i] = nums[i]
			}
		}
	}

	//return append(nums, nums...)
}

func main() {
	nums := []int{1, 3, 2, 1}
	//nums := []int{}
	c := getConcatenation(nums)
	fmt.Println(c)
}
