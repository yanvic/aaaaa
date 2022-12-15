package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func main() {
	fmt.Println("arquivo structs")
	//primeira forma
	var u usuario
	u.nome = "davi"
	u.idade = 21
	fmt.Println(u)
	//segunda forma
	usuario2 := usuario{"davi", 21}
	fmt.Println(usuario2)
	//terceira forma
	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}
