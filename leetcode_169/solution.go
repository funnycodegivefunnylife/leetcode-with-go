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

func majorityElementRequireO1Space(nums []int) int {
	var candidate, count int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
