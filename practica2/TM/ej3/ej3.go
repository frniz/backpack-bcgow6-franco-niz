package main

const (
	A = "A"
	B = "B"
	C = "C"
)

func calcularSalario(minutosTrabajados int, categoria string) (salario float64) {

	var horasMensual float64
	horasMensual = float64(minutosTrabajados) / 60.0
	switch categoria {
	case A:
		salario = (horasMensual * 3000)
		salario += salario * 0.5
	case B:
		salario = (horasMensual * 1500)
		salario += salario * 0.2
	case C:
		salario = (horasMensual * 1000)
	}

	return
}

func main() {

}
