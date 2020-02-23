package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[17952368] = "Pedro"
	aprovados[14679535] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[14679535])
	delete(aprovados, 14679535)
	fmt.Println(aprovados[14679535])
}
