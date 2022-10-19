package ej2

func OrdenarSliceDeEnteros(sliceDado []int) []int {
	slice := sliceDado
	for i, v1 := range slice {
		min := v1
		for j := i; j < len(slice); j++ {
			if slice[j] < min {
				aux := slice[j]
				slice[j] = min
				min = aux
			}
		}
		slice[i] = min
	}

	return slice
}
