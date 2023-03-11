package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/biancarosa/algo/sorting"
)

func main() {
	fmt.Println("Oi")
	tipo := os.Args[0]
	lista := os.Args[2:]
	fmt.Println(lista)

	// 2 coisas pra gente pensar melhor:
	// Como inferir o tipo sem o usuario tem que passar como argumento
	// Como fazer com que esse codigo dentro do if seja menor e mais reaproveitavel
	if tipo == "int" {
		listaDeInteiro := make([]int, len(lista))
		for i, s := range lista {
			var n int
			n, err := strconv.Atoi(s)
			if err != nil {
				panic("Argumento string invalido!")
			}
			listaDeInteiro[i] = n
		}
		ordenador := new(sorting.BubbleSort[int])
		listaOrdenada := ordenador.Sort(listaDeInteiro)
		fmt.Println(listaOrdenada)
	} else if tipo == "float64" {

	}

}
