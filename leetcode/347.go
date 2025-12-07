package leetcode

func topKFrequent(nums []int, k int) []int {
	// 1. Count frequency
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// 2. Bucket sort by frequency
	buckets := make([][]int, len(nums)+1)
	for num, f := range freq {
		buckets[f] = append(buckets[f], num)
	}

	// 3. Get the k elements with the highest frequency
	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		for _, num := range buckets[i] {
			res = append(res, num)
			if len(res) == k {
				break
			}
		}
	}

	return res
}
