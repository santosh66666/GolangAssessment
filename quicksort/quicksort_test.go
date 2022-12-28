package main

import (
    "testing"
)

func QuickSortTest(t * testing.T) {
	var arr = []int { 15, 3, 12, 6, -9, 9, 0 }
	quickSort(arr, 0, len(arr) - 1)
	t.Logf("Success")
}