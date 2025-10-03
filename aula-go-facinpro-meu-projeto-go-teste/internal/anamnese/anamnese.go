package anamnese

import (
	"fmt"
	"math"
)

func Anamnese () {
	nome := "Ana Flora"
	idade := 20
	peso := 61.35
	altura := 1.7
	imc := peso / math.Pow(altura, 2)
	fmt.Printf("Ficha Técnica: \n")
	fmt.Printf("Nome.....: %v\n", nome)
	fmt.Printf("Idade....: %v\n", idade)
	fmt.Printf("Peso.....: %v\n", peso)
	fmt.Printf("Altura...: %v\n", altura)
	fmt.Printf("IMC......: %v", imc)
}
