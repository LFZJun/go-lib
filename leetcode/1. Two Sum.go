package leetcode

//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	t := make(map[int]int)
	for i, v := range nums {
		if vv, ok := t[v]; ok {
			return []int{vv, i}
		} else {
			t[target-v] = i
		}
	}
	return []int{}
}
