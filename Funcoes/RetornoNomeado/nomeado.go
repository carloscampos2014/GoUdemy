package main

import "fmt"

func main() {
	s, su := operacoes(10, 40)
	fmt.Println(s, su)
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8}

	total := soma(numeros...)

	defer fmt.Println(total)

	x := func(num ...int) int {
		valor := 0
		for _, v := range num {
			valor += v
		}
		return valor
	}(numeros...)

	fmt.Println(x)

	fmt.Println(fibonati(40))
}

func operacoes(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func soma(num ...int) (total int) {
	for _, v := range num {
		total += v
	}
	return
}

func fibonati(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonati(posicao-2) + fibonati(posicao-1)
}
