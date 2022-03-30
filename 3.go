package main

import (
	"fmt"
	"path"
)

func main() {

	var dir, file string
	dir, file = path.Split("css/main.css")

	fmt.Println("dir:", dir)
	fmt.Println("file:", file)

}

//dir = diretório
//file = arquivo
//path package serve para trabalhar com o caminho das urls no formato strings
// dir é string e string também é
