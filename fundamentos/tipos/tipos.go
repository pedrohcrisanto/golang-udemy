package main 

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de il é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(49.99))
	fmt.Println(!bo)

	s1 := "OLá meu é Pedro"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Olá
	meu
	nome
	é
	Pedro
	`
	fmt.Println("O tamanho da string é", len(s2))

	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}