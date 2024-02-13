package searching

func (input *Input) LinearSearch(sItem int) int {
	for i, item := range *input {
		if item == sItem {
			return i
		}
	}
	return -1
}
