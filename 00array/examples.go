package array

import (
	"fmt"
)

// i/p []int{5, 10, 20, 15}
// o/p 20
func (input Input) FindPeakElement() []int {
	var result []int
	var ptr1, ptr2 int

	for index, val := range input.ItemsOfInt {
		ptr1 = index - 1
		ptr2 = index + 1

		if ptr1 < 0 {
			ptr1 = 0
		}

		if ptr2 >= len(input.ItemsOfInt) {
			ptr2 = len(input.ItemsOfInt) - 1
		}

		if input.ItemsOfInt[ptr1] < val && input.ItemsOfInt[ptr2] < val {
			result = append(result, val)
		}
	}

	return result
}

// i/p []int{1, 423, 6, 46, 34, 23, 13, 53, 4}
// o/p 1, 423
func (input Input) FindMinMax() []int {
	var min, max = input.ItemsOfInt[0], input.ItemsOfInt[0]

	for _, val := range input.ItemsOfInt {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	return []int{min, max}
}

// i/p []int{10, 20, 15, 4, 13}
// o/p []int{13, 4, 15, 20, 10}

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Pop() int {
	length := len(s.items) - 1
	val := s.items[length]
	s.items = s.items[:length]
	return val
}
func (input Input) ReverseArray() []int {
	result := []int{}
	stack := Stack{}

	for _, item := range input.ItemsOfInt {
		stack.Push(item)
	}

	for stack.Size() != 0 {
		result = append(result, stack.Pop())
	}

	return result
}

// i/p []int{10, 20, 15, 4, 13}
// o/p []int{4, 10, 13, 15, 20}
func (input Input) Sort() []int {

	for i := 0; i < len(input.ItemsOfInt); i++ {
		for j := i + 1; j < len(input.ItemsOfInt); j++ {
			if input.ItemsOfInt[i] > input.ItemsOfInt[j] {
				input.ItemsOfInt[i], input.ItemsOfInt[j] = input.ItemsOfInt[j], input.ItemsOfInt[i]
			}
		}
	}

	return input.ItemsOfInt
}

// i/p []int{10, 20, 15, 4, 13}, 2
// o/p []int{10, 15}
func (input Input) KthSmallestLargest(kth int) []int {
	input.Sort()

	return []int{input.ItemsOfInt[kth-1], input.ItemsOfInt[len(input.ItemsOfInt)-kth]}
}

// i/p []int{10, 20, 15, 4, 13}, 4
// o/p 1
func (input Input) FindXOccurrence(x int) int {
	inputMap := make(map[int]int)

	for i := 0; i < len(input.ItemsOfInt); i++ {
		inputMap[input.ItemsOfInt[i]]++
	}

	return inputMap[x]
}

func (input Input) DoSomething(x int) [][]int {

	result := make([][]int, x)

	for i := range result {
		start := i * x
		end := start + x
		if end > len(input.ItemsOfInt) {
			end = len(input.ItemsOfInt)
		}

		result[i] = input.ItemsOfInt[start:end]
	}

	return result
}

// i/p []int{0, 1, 2, 0, 1, 1, 2, 2}
// o/p []int{0, 0, 1, 1, 2, 2, 2}
func (input Input) Sort012() []int {
	count0, count1, count2 := 0, 0, 0
	result := []int{}

	// { 0: 2, 1: 3, 2:3 }
	for _, item := range input.ItemsOfInt {
		if item == 0 {
			count0++
		}

		if item == 1 {
			count1++
		}

		if item == 2 {
			count2++
		}
	}

	for i := 0; i < count0; i++ {
		result = append(result, 0)
	}

	for i := 0; i < count1; i++ {
		result = append(result, 1)
	}

	for i := 0; i < count2; i++ {
		result = append(result, 2)
	}

	return result
}

// i/p []int{10, 20, 15, 4, 13}, sum = 39
// o/p []int{1, 3}
func (input Input) SubArraySum(target int) []int {
	for i := 0; i < len(input.ItemsOfInt); i++ {
		currentSum := input.ItemsOfInt[i]
		for j := i + 1; j < len(input.ItemsOfInt); j++ {
			currentSum += input.ItemsOfInt[j]

			if currentSum == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// i/p []int{10, -20, -15, 4, 13}
// o/p []int{-20, -15, 10, 4, 13}
func (input Input) MoveNegElement() []int {
	for i := 0; i < len(input.ItemsOfInt); i++ {
		for j := i + 1; j < len(input.ItemsOfInt); j++ {
			if input.ItemsOfInt[i] > input.ItemsOfInt[j] && input.ItemsOfInt[j] < 0 {
				input.ItemsOfInt[i], input.ItemsOfInt[j] = input.ItemsOfInt[j], input.ItemsOfInt[i]
			}
		}
	}

	return input.ItemsOfInt
}

// i/p []int{10, 20, -15, 4, -4, 13}
// o/p [-15 -4 10 4 20 13]
func checkItemExist(arr []int, item int) bool {
	for _, val := range arr {
		if item == val {
			return true
		}
	}
	return false
}
func (input Input) UnionAndIntersection(arr []int) ([]int, []int) {
	union := []int{}
	intersection := []int{}

	unionMap := make(map[int]int)
	for i, item := range input.ItemsOfInt {
		unionMap[item] = i
	}

	for i, item := range arr {
		unionMap[item] = i
	}

	for key := range unionMap {
		union = append(union, key)
	}

	for i := 0; i < len(input.ItemsOfInt); i++ {
		if checkItemExist(arr, input.ItemsOfInt[i]) {
			intersection = append(intersection, input.ItemsOfInt[i])
		}
	}

	return union, intersection
}

// i/p []int{10, 20, 15, 4, 4, 13}
// o/p []int{13, 10, 20, 15, 4, 4}
func (input Input) RotateArray(x int) []int {

	last := input.ItemsOfInt[len(input.ItemsOfInt)-1]
	for i := len(input.ItemsOfInt) - 1; i >= 1; i-- {
		input.ItemsOfInt[i] = input.ItemsOfInt[i-1]
	}
	input.ItemsOfInt[0] = last

	return input.ItemsOfInt
}

// i/p []int{1, 2, 3, 4, 5, 6, 7}
// o/p 8
func (input Input) FindMissingNumber() int {
	arrSum := 0

	n := len(input.ItemsOfInt)
	sum := ((n + 2) * (n + 1)) / 2

	for _, item := range input.ItemsOfInt {
		arrSum += item
	}

	return sum - arrSum
}

// i/p []int{1, 2, 3, 4, 5, 6, 7, 9}, 13
// o/p 2
func (input Input) CountPairs(target int) int {
	count := 0

	compMap := make(map[int]int)

	for i, item := range input.ItemsOfInt {
		comp := target - item
		if _, ok := compMap[comp]; ok {
			count++
		}

		compMap[item] = i
	}

	return count
}

// i/p []int{1, 2, 3, 4, 5, 6, 7, 9, 2, 3}
// o/p []int{2, 3}
func (input Input) FindDuplicate() []int {
	result := []int{}
	inputMap := make(map[int]int)

	for _, item := range input.ItemsOfInt {
		inputMap[item]++
	}

	for k, v := range inputMap {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}

// i/p []int{1, 2, 3, 4, 5, 6, 7, 9, 2}, ]int{3, 5}, []int{3, 5}
// o/p []int{3, 5}
func (input Input) FindCommonItemIn3Array(arr1, arr2 []int) []int {
	result := []int{}

	for _, item := range input.ItemsOfInt {
		if checkItemExist(arr1, item) && checkItemExist(arr2, item) {
			result = append(result, item)
		}
	}

	return result
}

// i/p []int{1, 2, 3, 3, 4, 5, 6, 7, 9, 2}
// o/p 2
func (input Input) Find1stDuplicate() int {
	result := -1

	for i := 0; i < len(input.ItemsOfInt); i++ {
		if checkItemExist(input.ItemsOfInt[i+1:], input.ItemsOfInt[i]) {
			return input.ItemsOfInt[i]
		}
	}

	return result
}

// i/p []int{1, 2, 3, 3, 4, 5, 6, 7, 9, 2}
// o/p 1
func (input Input) Find1stNonDuplicate() int {
	result := -1

	for i := 0; i < len(input.ItemsOfInt); i++ {
		if !checkItemExist(input.ItemsOfInt[i+1:], input.ItemsOfInt[i]) {
			return input.ItemsOfInt[i]
		}
	}

	return result
}

// i/p []int{2, 3, 4, 1, 0}
// o/p 4
func check01Equal(arr [][]int) int {
	count := 0
	for _, items := range arr {
		itemMap := make(map[int]int)

		for _, val := range items {
			itemMap[val]++
		}

		tmp := []int{}
		for k, v := range itemMap {
			if k == 1 || k == 0 {
				tmp = append(tmp, v)
			}
		}

		if len(tmp) == 2 && tmp[0] == tmp[1] {
			count++
		}
	}

	return count
}
func (input Input) CountSubarrayWithEqual01() int {
	subArray := [][]int{}

	for i := 0; i < len(input.ItemsOfInt); i++ {
		for j := i; j < len(input.ItemsOfInt); j++ {
			sarr := []int{}
			for k := i; k <= j; k++ {
				sarr = append(sarr, input.ItemsOfInt[k])
			}
			subArray = append(subArray, sarr)
		}

	}

	fmt.Println(subArray)

	return check01Equal(subArray)
}

func (intput Input) TransposeMatrix() [][]int {
	matrix := intput.ItemsOfIntMatrix
	n := len(matrix)
	// for i := 0; i < n; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
	// 	}
	// }

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	tmp := []int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp = append(tmp, matrix[j][i])
		}
	}
	fmt.Println(tmp)

	return matrix
}

func (intput Input) RotateMatrix() [][]int {
	matrix := intput.ItemsOfIntMatrix
	m := len(matrix)
	n := len(matrix[0])
	row := 0
	col := 0
	var prev, curr int

	for row < m && col < n {
		if row+1 == m || col+1 == n {
			break
		}

		// store 1st ele of next row
		// this element will be replace by first element for current row
		prev = matrix[row+1][col]
		for i := col; i < n; i++ {
			curr = matrix[row][i]
			matrix[row][i] = prev
			prev = curr
		}
		row++

	}

	return matrix
}

func (intput Input) SortMatrix() [][]int {
	matrix := intput.ItemsOfIntMatrix
	m := len(matrix)
	n := len(matrix[0])
	arr := []int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr = append(arr, matrix[i][j])
		}
	}

	var array Input
	array.ItemsOfInt = arr
	array.Sort()

	k := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = array.ItemsOfInt[k]
			k++
		}
	}

	return matrix
}

func findMaxItem(arr []int) int {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
func (input Input) MaxItemFromRow() []int {
	matrix := input.ItemsOfIntMatrix
	maxItems := []int{}

	for i := 0; i < len(matrix); i++ {
		maxItems = append(maxItems, findMaxItem(matrix[i]))
	}

	return maxItems
}

func isSorted(arr []int) bool {

	left, right := 0, len(arr)-1

	//[1, 2,3]
	for left < right {
		if arr[left] > arr[left+1] {
			return false
		}
		left++
	}

	return true
}
func (input Input) CountSortedRows() int {
	matrix := input.ItemsOfIntMatrix
	count := 0

	for i := 0; i < len(matrix); i++ {
		if isSorted(matrix[i]) {
			count++
		}
	}

	return count
}

func itemExistInRow(arr1, arr2 []int) int {
	for _, item1 := range arr1 {
		for _, item2 := range arr2 {
			if item1 == item2 {
				return item1
			}
		}
	}
	return -1
}
func (input Input) CommonItemInRow() []int {
	result := []int{}
	matrix := input.ItemsOfIntMatrix
	firstRow := matrix[0]
	itemsMap := make(map[int]int)

	for i := 1; i < len(matrix); i++ {
		if item := itemExistInRow(firstRow, matrix[i]); item > 0 {
			itemsMap[item]++
		}
	}

	fmt.Println(itemsMap)

	return result
}

func is1Exists(arr []int) int {
	count1s := 0

	for _, item := range arr {
		if item == 1 {
			count1s++
		}
	}

	return count1s
}
func (input Input) RowWithMax1s() map[int]int {
	matrix := input.ItemsOfIntMatrix
	itemsMap := make(map[int]int)

	for i, row := range matrix {
		itemsMap[i] = is1Exists(row)
	}

	return itemsMap
}

// i/p {“geeksforgeeks”, “geeks”, “geek”, “geezer”}
// o/p gee
func (input Input) LargestCommonPrefix() string {
	var result string

	str1 := input.ItemsOfStr[0]
	for i := 0; i < len(input.ItemsOfStr); i++ {
		if str1[i] == input.ItemsOfStr[i][i] {
			result += string(str1[i])
		}
	}

	return result
}

func (input Input) DiamondTraversal(arr [][]int) []int {
	sum := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i+j == 1 || i+j == 3 {
				sum = append(sum, arr[i][j])
			}
		}
	}

	return sum
}

// var x = [][]int{
// {2, 3, 4},
// {5, 7, 9},
// {6, 8, 10}}
