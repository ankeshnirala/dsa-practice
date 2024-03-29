package main

import (
	linkedlistrevision01 "dsa-practice/01linkedlist-revision-01"
)

func main() {
	// var array array.Input

	// array.ItemsOfInt = []int{10, 20, 15, 4, 4, 13}

	// start := time.Now()
	// fmt.Println(array.QuickSort(0, len(array.ItemsOfInt)-1))
	// fmt.Println(time.Since(start))

	// start := time.Now()
	// sort.Ints(array.ItemsOfInt)
	// fmt.Println(array.ItemsOfInt)
	// fmt.Println(time.Since(start))
	// fmt.Println("FindPeakElement: ", array.FindPeakElement())
	// fmt.Println("FindMinMax: ", array.FindMinMax())
	// fmt.Println("ReverseArray: ", array.ReverseArray())
	// fmt.Println("SortArray: ", array.Sort())
	// fmt.Println("KthSmallestLargest:", array.KthSmallestLargest(3))
	// fmt.Println("FindXOccurrence:", array.FindXOccurrence(4))
	// fmt.Println("DoSomething", array.DoSomething(3))

	// array.ItemsOfInt = []int{0, 1, 2, 0, 1, 1, 2, 2}
	// fmt.Println("Sort012:", array.Sort012())
	// array.ItemsOfInt = []int{10, 20, 15, 4, 4, 13}
	// fmt.Println("SubArraySum:", array.SubArraySum(39))
	// array.ItemsOfInt = []int{10, 20, -15, 4, -4, 13}
	// fmt.Println("MoveNegElement: ", array.MoveNegElement())
	// fmt.Printf("UnionAndIntersection:")
	// fmt.Println(array.UnionAndIntersection([]int{4, 30}))

	// array.ItemsOfInt = []int{10, 20, 15, 4, 4, 13}
	// fmt.Println("RotateArray:", array.RotateArray(2))

	// array.ItemsOfInt = []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println("FindMissingNumber:", array.FindMissingNumber())

	// array.ItemsOfInt = []int{2, 3, 3, 4, 5, 6, 7, 9, 2}
	// fmt.Println("CountPairs: ", array.CountPairs(13))
	// fmt.Println("FindDuplicate: ", array.FindDuplicate())
	// fmt.Println("FindCommonItemIn3Array:", array.FindCommonItemIn3Array([]int{3, 5}, []int{3, 5}))
	// fmt.Println("Find1stDuplicate:", array.Find1stDuplicate())
	// fmt.Println("Find1stNonDuplicate:", array.Find1stNonDuplicate())

	// array.ItemsOfInt = []int{2, 3, 4, 1, 0}
	// fmt.Println("CountSubarrayWithEqual01:", array.CountSubarrayWithEqual01())
	// fmt.Println("DiamondTraversal: ", array.DiamondTraversal([][]int{{2, 3, 4}, {5, 7, 9}, {6, 8, 10}}))

	// array.ItemsOfIntMatrix = [][]int{{2, 3, 4}, {5, 7, 9}, {6, 8, 10}}
	// fmt.Println("TransposeMatrix: ", array.TransposeMatrix())
	// fmt.Println("RotateMatrix: ", array.RotateMatrix())
	// fmt.Println("SortMatrix:", array.SortMatrix())
	// fmt.Println("MaxItemFromRow:", array.MaxItemFromRow())
	// array.ItemsOfIntMatrix = [][]int{{2, 1, 4}, {4, 1, 1}, {9, 120, 10}}
	// fmt.Println("CountSortedRows:", array.CountSortedRows())
	// fmt.Println("CommonItemInRow:", array.CommonItemInRow())
	// fmt.Println("RowWithMax1s:", array.RowWithMax1s())

	// array.ItemsOfStr = []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	// array.ItemsOfStr = []string{"apple", "ape", "april"}
	// fmt.Println("LargestCommonPrefix: ", array.LargestCommonPrefix())

	ll := linkedlistrevision01.LinkedList{}

	ll.Insert("10")
	ll.Insert("20")
	ll.Insert("30")
	ll.Insert("40")

	// fmt.Println(ll)
	ll.Display()
}
