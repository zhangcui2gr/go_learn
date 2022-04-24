package main

import "fmt"

func swap(num1 *int, num2 *int) {
	t := *num1
	*num1 = *num2
	*num2 = t
}

func bubble_sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

func quick_sort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	for i < j {
		for i < j && arr[j] >= arr[l] {
			j--
		}
		for i < j && arr[i] <= arr[l] {
			i++
		}
		swap(&arr[i], &arr[j])
	}
	swap(&arr[l], &arr[i])
	quick_sort(arr, l, i)
	quick_sort(arr, i+1, r)

}

func main() {
	arr := []int{3, 8, 1, 2, 0, 5, 6, 4, 7}
	fmt.Println(bubble_sort(arr))

	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
