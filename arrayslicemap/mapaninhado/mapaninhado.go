package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 123123.78,
			"Guga Pereira":   123123.56,
		},
		"J": {
			"José João": 11345.45,
		},
		"P": {
			"Pedro Junior": 12312.2,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
