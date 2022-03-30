/*greeter
os package , fala com diversos sistemas operacionais
args = argumentos
slice = slice pode guardar multiplos valores
go é indexado por zero
então o primeiro É ZERO, O SEGUNDO É 1
COMO USÁ-LAS?
go run cria caminhos temporários
go build cria um arquivo fixo
nesse caso rodamos o seguintes comandos : go build10.go , ./10   e  ./10 hi hellow hey
também usamos o lenght ou len, para mostrar o comprimento desses dados que foram inseridos nesse slide
buildamos novamente, executamos com as frases nessa slice e veremos que agora temos 4 argumentos*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2st argument:", os.Args[2])
	fmt.Println("3st argument:", os.Args[3])

	fmt.Println("NUMBER OF ITENS os.Args: ", len(os.Args))

}
