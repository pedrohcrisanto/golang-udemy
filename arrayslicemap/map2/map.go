package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose João":      123412.45,
		"Gabriela Silva": 123152.52,
		"Pedro Junior":   1231.1,
	}

	funcsESalarios["Rafael Filho"] = 12312.2
	delete(funcsESalarios, "Inesxistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
