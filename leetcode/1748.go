package leetcode

func sumOfUnique(nums []int) int {
	mp := map[int]int{}

	for _, v := range nums {
		mp[v]++
	}

	sum := 0
	for k, v := range mp {
		if v < 2 {
			sum += k
		}
	}
	return sum
}

func sumOfUniqueV2(nums []int) int {
	mp := map[int]int{}
	sum := 0

	for _, v := range nums {
		mp[v]++
		if mp[v] == 1 {
			sum += v
		} else if mp[v] == 2 {
			sum -= v
		}
	}

	return sum
}
