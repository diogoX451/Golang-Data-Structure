package ordenation

func QuickSort(array *[]int, left int, right int) {
	if left < right {
		pivot := partition(array, left, right)
		QuickSort(array, left, pivot-1)
		QuickSort(array, pivot+1, right)
	}
}

func partition(array *[]int, left int, rigth int) int {
	pivot := (*array)[left]
	pointer := left

	for i := left + 1; i <= rigth; i++ {
		if (*array)[i] <= pivot {
			pointer++
			replacements(*array, i, pointer)
		}
	}

	replacements(*array, left, pointer)

	return pointer
}

func replacements(array []int, i int, j int) {
	aux := array[i]
	array[i] = array[j]
	array[j] = aux
}
