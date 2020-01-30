package main

import "fmt"

func fatorial(n uint) uint {
	if n == 0 {
		return 1
	}
	anterior := fatorial(n - 1)
	return n * anterior

}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

}
