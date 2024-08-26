package leetcode_268

func missingNumber(nums []int) int {
	size := len(nums)
	expectSum := size * (size + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return expectSum - sum
}
