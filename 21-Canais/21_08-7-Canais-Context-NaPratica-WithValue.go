package main

import (
	"context"
	"fmt"
)

/*
=> Context => serve pra trocar mensagens entre GO Routines
- Só pra ter uma idéia geral:
- Se a gente lança 500 goroutines pra fazer uma tarefa, e cancelamos a tarefa no meio do caminho, como fazemos pra matar as goroutines?
- Documentação: https://golang.org/pkg/context/
- Aos aventureiros: https://blog.golang.org/context
- Destaques:
    - ctx := context.Background (base do context)
    - ctx, cancel = context.WithCancel(context.Background) (cancela Go Routine)
    - goroutine: select case ←ctx.Done(): return; default: continua trabalhando.
    - check ctx.Err();
    - Tambem tem:
			a) WithDeadline(horario que se espera pra que se finalize uma execucao) /
			b) Timeout (quantidade de min/horas para executar algo)
- Exemplos (Todd):
    - Analisando:
        - Background: https://play.golang.org/p/cByXyrxXUf
        - WithCancel: https://play.golang.org/p/XOknf0aSpx
        - Função Cancel: https://play.golang.org/p/UzQxxhn_fm
    - Exemplos práticos:
        - func WithCancel: https://play.golang.org/p/Lmbyn7bO7e
        - func WithCancel: https://play.golang.org/p/wvGmvMzIMW
        - func WithDeadline: https://play.golang.org/p/Q6mVdQqYTt
        - func WithTimeout: https://play.golang.org/p/OuES9sP_yX
        - func WithValue: https://play.golang.org/p/8JDCGk1K4P
*/

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go")
	foo(ctx, "language")

	ctx = context.WithValue(ctx, "dog", "Gaston")
	foo(ctx, "dog")

	foo(ctx, "color")
}

func foo(ctx context.Context, s string) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", s)
}