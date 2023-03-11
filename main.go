package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/biancarosa/algo/sorting"
)

func main() {
	fmt.Println("Oi")
	lista := os.Args[1:]
	fmt.Println(lista)

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
}
