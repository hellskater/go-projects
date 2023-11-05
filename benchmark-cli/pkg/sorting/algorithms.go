package sorting

func BubbleSort(elements []int) {
	for i := 0; i < len(elements)-1; i++ {
		if elements[i] > elements[i+1] {
			elements[i], elements[i+1] = elements[i+1], elements[i]
		}
	}
}

func MergeSort(elements []int) {
	if len(elements) < 2 {
		return
	}

	mid := len(elements) / 2
	MergeSort(elements[:mid])
	MergeSort(elements[mid:])

	sorted := make([]int, len(elements))
	i, j := 0, mid

	for k := 0; k < len(elements); k++ {
		if i < mid && (j >= len(elements) || elements[i] < elements[j]) {
			sorted[k] = elements[i]
			i++
		} else {
			sorted[k] = elements[j]
			j++
		}
	}

	copy(elements, sorted)
}

func QuickSort(elements []int) {
	if len(elements) < 2 {
		return
	}

	pivot := elements[0]
	i, j := 1, len(elements)-1

	for i <= j {
		for i <= j && elements[i] < pivot {
			i++
		}

		for i <= j && elements[j] >= pivot {
			j--
		}

		if i <= j {
			elements[i], elements[j] = elements[j], elements[i]
			i++
			j--
		}
	}

	elements[0], elements[j] = elements[j], elements[0]

	QuickSort(elements[:j])
	QuickSort(elements[j+1:])
}