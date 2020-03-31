package math

func Sum(nums []int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func MinIn(nums []int) int {
	result := nums[0]
	for _, n := range nums {
		if n <= result {
			result = n
		}
	}
	return result
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func MaxIn(nums []int) int {
	result := 0
	for _, n := range nums {
		if n >= result {
			result = n
		}
	}
	return result
}
