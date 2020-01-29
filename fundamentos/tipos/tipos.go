package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("O tipo da variável i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais (float32, float64)
	var x float32 = 49.99 // ou var x = float32(49.99)
	fmt.Println("O tipo da variável x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal do 49.99 é:", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Rodrigo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é:", len(s1))

	// string com múltiplas linhas
	s2 := ` Olá
	meu
	nome
	é
	Rodrigo`
	fmt.Println("O tamanho da string é:", len(s2))

	//char????
	// var x char = 'b'   -> não existe
	char := 'a'
	fmt.Println("O tipo do char é:", reflect.TypeOf(char))
}