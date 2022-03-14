package main

import (
	"log"
	"os"

	"dadosweb/app"
)

func main() {
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatalln(erro.Error())
	}
}
