package leetcode_169

func majorityElement(nums []int) int {
	var dict = make(map[int]int)

	for _, num := range nums {
		if _, ok := dict[num]; ok {
			dict[num]++
		} else {
			dict[num] = 1
		}
	}

	var maxNum, maxCount int
	for num, count := range dict {
		if count > maxCount {
			maxCount = count
			maxNum = num
		}
	}

	return maxNum
}
