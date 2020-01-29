package main

import "fmt"

func main() {
	i := 1 //ocupa 64 bits ou 8 bytes

	// Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i // pegando endereço da variável i
	*p++   // desreferenciando (pegando o valor)
	i++
	fmt.Println(p, *p, i, &i)

}
