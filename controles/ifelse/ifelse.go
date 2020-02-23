package main

import (
	"fmt"
)

func imprirmirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprirmirResultado(7.3)
	imprirmirResultado(5.1)
}
