package main

import "fmt"

func search(array []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := array[mid]
		if num == target {
			return mid
		}
		if num < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(search(a, 4))
}
