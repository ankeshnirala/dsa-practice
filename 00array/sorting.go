package array

func Partition(arr []int, lb, ub int) int {
	pivot := lb
	start, end := lb, ub

	for start < end {
		for arr[start] <= arr[pivot] {
			start++
		}

		for arr[end] > arr[pivot] {
			end--
		}

		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}

	arr[lb], arr[end] = arr[end], arr[lb]

	return end
}

func (input *Input) QuickSort(lb, ub int) []int {
	if lb < ub {
		loc := Partition(input.ItemsOfInt, lb, ub)
		input.QuickSort(lb, loc-1)
		input.QuickSort(loc+1, ub)
	}

	return input.ItemsOfInt
}
