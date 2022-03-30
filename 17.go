//uso da underline para ignorar o resto dos arquivos

package main

import (
	"fmt"
	"path"
)

func main() {

	dir, _ := path.Split("Secret/file.txt")
	fmt.Println(dir)

	//agora printando o file
	dir, file := path.Split("Secret/file.txt")
	fmt.Println(dir, file)
}
