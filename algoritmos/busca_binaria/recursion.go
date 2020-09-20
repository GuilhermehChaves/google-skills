package main

import "fmt"

func main() {
	numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(search(numbers, 30))
}

func search(a []int, element int) int {
	start := 0
	end := len(a) - 1
	return binary(a, element, start, end)
}

func binary(a []int, element, start, end int) int {
	mid := (start + end) / 2
	currentSize := end - start

	if a[mid] == element {
		return mid
	}

	if currentSize == 0 && a[mid] != element {
		return -1
	}

	if a[mid] > element {
		return binary(a, element, start, mid-1)
	}

	// if a[mid] < element
	return binary(a, element, mid+1, end)
}
