package main

import (
	"fmt"
	"lc/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Inicialização")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}