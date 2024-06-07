package main

import "fmt"

func BinarySearch(arr []int, target int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 3
	fmt.Println("Is target in array?", BinarySearch(arr, target))
}
