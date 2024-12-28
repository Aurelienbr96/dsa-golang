package binarysearch

func calcmid(low int, max int) int {
	return (low + max) / 2
}

func BinarySearch(arr []int, val int) int {

	max := len(arr)
	low := 0

	for low != max {
		mid := calcmid(low, max)
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			max = mid
		} else {
			low = mid + 1
		}
	}

	return -1
}
