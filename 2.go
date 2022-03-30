package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.NumCPU() + 1)
}

//NumCPU = retorna o número de processadores disponível no sistema operacional
// se for adicionado +1 após o parenteses, aumenta este número
