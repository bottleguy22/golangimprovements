package main

import "fmt"

func main() {
	speed := 100 //inteiro
	force := 2.5 //flutuante

	speed = speed * int(force)
	fmt.Println(speed)
	fmt.Println(force, int(force))
}

//você não pode usar tipos diferentes juntos
//a solução seria converter uma das duas e deixá-las iguais
//antes de usar o force, eu transformei em inteiro

//ele está arredondando para 2, como posso deixar em 2,5?


package main

import "fmt"

func main() {
	speed := 100
	force := 2.5

	speed = int(float64(speed) * (force))
	fmt.Println(speed)
}
//Eu posso englobar tudo transformando todos em inteiros
//virando float64(100) = 100.0 -> float 64
//100.0 * 2.5 = 250.0
