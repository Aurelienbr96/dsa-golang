package bubble_sort

func Sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for a := 1; a < len(arr); a++ {
			if arr[a-1] > arr[a] {
				temp := arr[a-1]
				arr[a-1] = arr[a]
				arr[a] = temp
			}
		}
	}
}
