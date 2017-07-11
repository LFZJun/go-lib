package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(twoSum([]int{1, 3, 2}, 1))
}

func Test2(t *testing.T) {
	result := addTwoNumbers(
		&listNode{5, nil},
		&listNode{5, nil},
	)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func Test3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}