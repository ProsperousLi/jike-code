func plusOne(digits []int) []int {
	size := len(digits)

	digits[size-1] += 1

	if digits[size-1] == 10 {
		digits[size-1] = 0
		for i := size - 2; i >= 0; i-- {
			if digits[i] < 9 {
				digits[i] += 1
				break
			}

			digits[i] = 0
		}
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits[0:]...)
	}

	return digits
}
