package main

func countInversions(arr []int32) int64 {
	a := int64(0)
	var shifts = &a
	_ = mergeSort(arr, shifts)

	return *shifts
}

func insertionSort(arr []int32) int64 {

	a := int64(0)
	var shifts = &a
	_ = mergeSort(arr, shifts)

	return *shifts
}

func mergeSort(arr []int32, shifts *int64) []int32 {
	arrLength := int32(len(arr))

	if arrLength == 1 {
		return arr
	}

	arrMiddle := arrLength / 2
	var left = make([]int32, arrMiddle)
	var right = make([]int32, arrLength-arrMiddle)

	for i := int32(0); i < arrLength; i++ {
		if i < arrMiddle {
			left[i] = arr[i]
		} else {
			right[i-arrMiddle] = arr[i]
		}
	}

	return merge(mergeSort(left, shifts), mergeSort(right, shifts), shifts)
}

func merge(left, right []int32, shifts *int64) []int32 {
	result := make([]int32, len(left)+len(right))

	i := int32(0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
			*shifts += int64(len(left))
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func main() {}
