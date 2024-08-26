package main

import (
	"api/src/config"
	"api/src/router"
	. "fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	r := router.Gerar()

	Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(Sprintf(":%d", config.Porta), r))
}
