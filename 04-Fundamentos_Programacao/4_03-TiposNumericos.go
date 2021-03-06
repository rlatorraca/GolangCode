package main

/*
 int vs. float: Números inteiros vs. números com frações.
- golang.org/ref/spec → numeric types
- Integers:
    - Números inteiros
    - int & uint → “implementation-specific sizes”
    - Todos os tipos numéricos são distintos, exceto:
        - byte = uint8
        - rune = int32 (UTF8)
        (O código fonte da linguagem Go é sempre em UTF-8).
    - Tipos são únicos
        - Go é uma linguagem estática
        - int e int32 não são a mesma coisa
        - Para "misturá-los" é necessário conversão
    - Regra geral: use somente int
- Floating point:
    - Números racionais ou reais
    - Regra geral: use somente float64
- Na prática:
    - Defaults com :=
    - Tipagem com var
    - Dá pra colocar número com vírgula em tipo int?
    - Overflow
    - Go Playground: https://play.golang.org/p/dt2x1ies5b
- “implementation-specific sizes”? Runtime package. Word.
    - GOOS
    - GORUNTIME
    - https://play.golang.org/p/1vp5DImIMM
 */

import (
	"fmt"
	"runtime"
)
var xx = 10
var yy = 10.0

func main() {
	a := "e"
	b := "é"
	c := "香"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a) // slice of bytes
	e := []byte(b) // slice of bytes
	f := []byte(c) // slice of bytes

	fmt.Printf("%v, %v, %v\n", d, e, f)
	fmt.Println()
	fmt.Printf("%v, %T\n", xx, xx)
	fmt.Printf("%v, %T\n", yy, yy)
	fmt.Println()
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)


}




/*
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
 */