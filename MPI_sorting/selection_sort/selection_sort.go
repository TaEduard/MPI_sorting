package selection_sort

func Sort(items []int) {
	n := len(items)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
