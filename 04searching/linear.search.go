package searching

func LinearSearch(input []int, sItem int) int {
	for i, item := range input {
		if item == sItem {
			return i
		}
	}
	return -1
}
