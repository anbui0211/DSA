package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	mapNum := map[int]int{}
	for _, v := range nums1 {
		mapNum[v]++
	}

	result := make([]int, 0)
	for _, v := range nums2 {
		if value, ok := mapNum[v]; ok {
			if value > 0 {
				result = append(result, v)
				mapNum[v] = -1
			}
		}
	}
	return result
}

func intersectionV2(nums1 []int, nums2 []int) []int {
	seen := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		seen[nums1[i]] = true
	}

	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if seen[nums2[i]] {
			res = append(res, nums2[i])
			seen[nums2[i]] = false // đảm bảo unique
		}
	}
	return res

	// Way-2: using for range
	// for _, v := range nums1 {
	// 	seen[v] = true
	// }

	// res := make([]int, 0)
	// for _, v := range nums2 {
	// 	if seen[v] {
	// 		res = append(res, v)
	// 		seen[v] = false
	// 	}
	// }
	// return res
}
