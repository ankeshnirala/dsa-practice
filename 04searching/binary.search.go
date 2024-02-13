package searching

func (input *Input) BinarySearch(sItem int) int {
	beg, end := 0, len(*input)-1

	for beg <= end {
		mid := (beg + end) / 2

		if (*input)[mid] == sItem {
			return mid
		} else if sItem < (*input)[mid] {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}

	return -1
}
