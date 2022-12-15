package main

import (
	"errors"
	"fmt"
)

// tipos de variaveis
func main() {
	var variavel string = "variavel 1"
	variavel2 := "variavel 2"
	fmt.Println(variavel)
	fmt.Println(variavel2)

	var (
		variavel3 string = "var3"
		variavel4 string = "var4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante1"
	fmt.Println(constante1)
	//inverter variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}

// tipos de numeros
func numero() {
	var numero int = 100
	fmt.Println(numero)
	// uint -(apenas numero sem sinal) int sem sinal
	var numero2 uint32 = 10000
	fmt.Println(numero2)
	//alias
	//int32 = RUNE
	var numero3 rune = 123
	fmt.Println(numero3)
	//BYTE = UINT8
	var numero4 byte = 124
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.50
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 13.50
	fmt.Println(numeroReal2)

	var numeroReal3 = 23.50
	fmt.Println(numeroReal3)

	var palavra string = "texto"
	fmt.Println(palavra)

	palavra2 := "texto2"
	fmt.Println(palavra2)

	//booleano
	var booleano1 bool
	fmt.Println(booleano1)
	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto da função 01")
	fmt.Println(resultado)

	_, resultadoSubtracao := calculosMatematicos(10, 20)
	fmt.Println(resultadoSubtracao)
}
func main() {
	//aritmeticos
	soma := 1 + 5
	subtracao := 1 - 5
	divisao := 10 / 4
	multiplicacao := 10 * 20
	restoDaDivisao := 10 % 3
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//operadores logicos
	fmt.Println(" ")
	verdadeiro, falso := true, false
	//todos verdadeiros
	fmt.Println(verdadeiro && falso)
	//todos falsos
	fmt.Println(verdadeiro || falso)
	//inverter
	fmt.Println(!verdadeiro)

	//operadores unarios
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3
}
