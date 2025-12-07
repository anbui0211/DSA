package leetcode

func twoSum(nums []int, target int) []int {
	mp := map[int]int{}

	for i, v := range nums {
		complement := target - v
		if index, ok := mp[complement]; ok {
			return []int{index, i}
		}
		mp[v] = i
	}

	return nil
}
