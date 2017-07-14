package leetcode

//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	start := -1
	longest := 0
	set := [256]int{}
	for i := range set {
		set[i] = -1
	}
	for i, r := range s {
		if v := set[r]; v > start {
			start = v
		} else {
			length := i - start
			if length > longest {
				longest = length
			}
		}
		set[r] = i
	}
	return longest
}
