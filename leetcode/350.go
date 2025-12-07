package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	mapNum := map[int]int{}
	for _, v := range nums1 {
		mapNum[v]++
	}

	result := make([]int, 0)
	for _, v := range nums2 {
		if value, ok := mapNum[v]; ok {
			if value > 0 {
				result = append(result, v)
				mapNum[v]--
			}
		}
	}

	return result
}

func intersectV2(nums1 []int, nums2 []int) []int {
	freq := make(map[int]int)
	for _, v := range nums1 {
		freq[v]++
	}

	res := make([]int, 0)
	for _, v := range nums2 {
		if freq[v] > 0 {
			res = append(res, v)
			freq[v]--
		}
	}

	return res
}
