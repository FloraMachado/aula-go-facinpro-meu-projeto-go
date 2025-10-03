// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
	"meu-projeto-go/internal/anamnese"
	"meu-projeto-go/internal/fibonacci"
	"meu-projeto-go/internal/hello"
	"meu-projeto-go/internal/saudacao"
)

// Função principal do programa
func main() {
	// Mensagem inicial da aplicação
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")

	// Chamada para função de saudação
	hello.SayHello()
	fmt.Println("-----------------------------")

	// Demonstração: cálculo do 10º número de Fibonacci
	n := 10
	// Chama a função Fibonacci do pacote fibonacci
	// fibonacci // importado acima
	// Fibonacci(n) // retorna o n-ésimo número da sequência
	// := é usado para declarar e inicializar a variável
	valor := fibonacci.Fibonacci(n)
	// Imprime o resultado com formatação
	fmt.Printf("F(%d) = %d\n", n, valor)

	// Demonstração: imprimir a sequência completa até n
	fibonacci.PrintSequence(n)

	fmt.Println("-----------------------------")
	anamnese.Anamnese()

	fmt.Println("\n\n-----------------------------")
	fmt.Println("Saudação:")

	mensagem := saudacao.Saudacao("Flora")
	fmt.Printf(mensagem)

	fmt.Println("\n\n-----------------------------")
	fmt.Println("Função anônima:")
	soma := func (n1, n2 int) int {
		return	n1 + n2
	}
	n1 := 2
	n2 := 3
	fmt.Printf("A soma de %v + %v = %v", n1, n2, soma(n1, n2))
}
