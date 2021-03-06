package main

import "fmt"

/*
- v, ok := ←chan
- Se receber valor: v, true
- Canal fechado, nada, etc.: zero v, false
- Agora vamos resolver o problema do exercício anterior usando comma ok.
- Código: https://github.com/ellenkorbes/aprend...
 */

//        - Chans par, ímpar, quit
//        - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
//        - Func receive é um select entre os três canais, encerra no quit

func main() {

	par := make(chan int)
	ímpar := make(chan int)
	quit := make(chan bool)

	go sendNúmerosY(par, ímpar, quit)

	receiverY(par, ímpar, quit)
}

func sendNúmerosY(par, ímpar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			ímpar <- i
		}
	}
	close(par)
	close(ímpar)
	quit <- true
}

func receiverY(par, ímpar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-ímpar:
			fmt.Println("O número", v, "é ímpar.")

		// Se recebe TRUE, encerra a funcao
		// Serve para evitar a impressao de numeros (vazios/seros)
		// e valores padroes do canal vazio
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra! Ó:", v)
			} else {
				fmt.Println("Encerrando. Recebemos:", v)
			}
			return
		}
	}
}

