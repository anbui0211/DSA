package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapArr := map[rune]int{}

	for _, char := range s {
		mapArr[char]++
	}

	for _, char := range t {
		mapArr[char]--
	}

	for _, v := range mapArr {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagramV2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := make(map[rune]int)

	for _, c := range s {
		cnt[c]++
	}

	for _, c := range t {
		cnt[c]--
		if cnt[c] < 0 {
			return false
		}
	}
	return true
}

func isAnagramV3(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var cnt [26]int

	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
