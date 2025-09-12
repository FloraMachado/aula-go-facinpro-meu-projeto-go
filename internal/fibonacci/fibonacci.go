package fibonacci

import "fmt"

func Fibonacci(n int) int {
	if n < 0 {
		panic("n não pode ser negativo")
	}
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func Sequence(n int) []int {
	if n < 0 {
		panic("n não pode ser negativo")
	}
	if n == 0 {
		return []int{0}
	}
	seq := []int{0, 1}
	for i := 2; i <= n; i++ {
		seq = append(seq, seq[i-1]+seq[i-2])
	}
	return seq
}

func PrintSequence(n int) {
	fmt.Printf("Sequência de Fibonacci até F(%d): %v\n", n, Sequence(n))
}
