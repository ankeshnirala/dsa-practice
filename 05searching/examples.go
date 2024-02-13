package searching

func (input Input) CountTargetOccurence(target int) int {
	count := 0

	for _, item := range input {
		if item == target {
			count++
		}
	}

	return count
}

func (input Input) FindMaxMin() (int, int) {
	max, min := input[0], input[0]

	for _, item := range input {
		if item < min {
			min = item
		}

		if item > max {
			max = item
		}
	}

	return min, max
}

func (input Input) TwoSum(target int) []int {
	i := 0
	for j := i + 1; j < len(input); j++ {
		if input[i]+input[j] == target {
			return []int{i, j}
		} else {
			i++
		}
	}

	return []int{-1, -1}
}
