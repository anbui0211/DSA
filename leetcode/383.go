package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	magazineMap := map[rune]int{}

	for _, ch := range magazine {
		magazineMap[ch]++
	}

	for _, ch := range ransomNote {
		if magazineMap[ch] == 0 {
			return false
		}
		magazineMap[ch]--
	}

	return true
}

func canConstructV2(ransomNote string, magazine string) bool {
	mp := make([]int, 26)
	for _, ch := range magazine {
		//ex: ch = 'a' -> int(ch) = 97
		// -> int(ch)-int('a') = 0 -> mp[0]
		// ->  mp[0]++ -> mp[0] = 1
		mp[ch-'a']++
	}

	for _, ch := range ransomNote {
		if mp[ch-'a'] == 0 {
			return false
		}
		mp[ch-'a']--
	}

	return true
}
