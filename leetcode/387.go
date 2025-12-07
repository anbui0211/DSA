package leetcode

// Way-1: using map
func firstUniqChar(s string) int {
	mp := make(map[rune]int)
	for _, ch := range s {
		mp[ch]++
	}

	for i, ch := range s {
		if mp[ch] == 1 {
			return i
		}
	}
	return -1
}

// Way-2: using array
func firstUniqCharV2(s string) int {
	arr := [26]int{}

	for _, ch := range s {
		arr[ch-'a']++
	}

	for i, ch := range s {
		if arr[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
