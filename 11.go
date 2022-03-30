/*formas idiomáticas são as melhores formas de codar algo, excesso de verbos causará confusão na leitura
aqui usamos o package buffer, ele é bem popular para usos em que o comprimento (LEN) da string é muito grande
b é apenas a abreviação de buffer,
n é o número de bytes lidos
off é offseats, uma espécie de variável do buffer*/
package main

func Read(b *Buffer, p []byte) (n int, err error) {
	if b.empty() {
		b.Reset()
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	return n, nil
}

/*abreviações comuns em GOLANG:
var s string // string
var i int // index
var msg string // message
var v string // value
var val string // value
var num int // number
var fv string // flag value
var err error // error value
var args []string // arguments
var seen bool // has seen?
var parsed bool // parsing ok?
var buf []byte // buffer
var off int // offset
var op int // operation
var opRead int // read operation
var l int // length
var n int // number or number of
var m int // another number
var c int // capacity
var c int // character
var a int // array
var r rune // rune
var sep string // separator
var src int // source
var dst int // destination
var b byte // byte
var w io.Writer // writer
var r io.Reader // reader
var pos int // position
var b []byte // buffer
var buf []byte // buffer*/
