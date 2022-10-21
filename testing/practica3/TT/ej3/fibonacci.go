package ej3

func Fibonacci(large int) (serie []int) {
	if large < 1 {
		return
	}
	if large == 1 {
		serie = append(serie, 0)
		return
	}

	serie = append(serie, 0, 1)

	if large > 2 {
		for i := 2; i < large; i++ {
			serie = append(serie, serie[i-1]+serie[i-2])
		}
	}

	return
}
