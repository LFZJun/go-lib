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
	var start, longest int
	chars := []rune(s)
	set := make(map[rune]int)
	for i, r := range chars {
		if v, ok := set[r]; ok && v >= start {
			start = v + 1
		} else {
			length := i - start + 1
			if length >= longest {
				longest = length
			}
		}
		set[r] = i
	}
	return longest
}
