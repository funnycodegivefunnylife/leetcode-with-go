package leetcode_2044

func orOperation(first int, second int) int {
	return first | second
}

// paramemeter is a function and array
func reduce(f func(int, int) int, arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = f(result, arr[i])
	}
	return result
}

// calculate the xor of all elements in the array which are selected by the bitmask
func calculateXorByBitmask(nums []int, bitMask int) int {

	size := len(nums)
	orResult := 0
	for i := 0; i < size; i++ {
		// if the bitMask is 1, then the bit is 1
		if (bitMask>>i)&1 == 1 {
			orResult |= nums[i]
		}
	}

	return orResult
}

func countMaxOrSubsets(nums []int) int {

	maxOr := reduce(orOperation, nums)

	count := 0

	// 1 << len(nums) is equal to 2^len(nums)
	for bitMask := 1; bitMask < 1<<len(nums); bitMask++ {
		if calculateXorByBitmask(nums, bitMask) == maxOr {
			count++
		}
	}

	return count
}
