package main

import (
	"fmt"
	"github.com/yanvic/aaaaa/struct"
)
type Main struct {
	none  usuario.
	idade int
}
// tipos de variaveis
func variaveis() {
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
