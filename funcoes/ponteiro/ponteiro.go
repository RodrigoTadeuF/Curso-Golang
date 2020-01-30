package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) { // função não pura por mexer em variáveis que estão fora dela
	//revisão * é usado para desreferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor ou seja a função operará por uma "cópia desse valor"
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variavel
	inc2(&n) // por referência
	fmt.Println(n)
}
