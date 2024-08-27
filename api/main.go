package main

import (
	"api/src/config"
	"api/src/router"
	. "fmt"
	"log"
	"net/http"
)

//func init() {
//	chave := make([]byte, 64)
//	if _, err := rand.Read(chave); err != nil {
//		log.Fatal(err)
//	}
//
//	stringBase64 := base64.StdEncoding.EncodeToString(chave)
//	println(stringBase64)
//}

func main() {
	config.Carregar()

	r := router.Gerar()

	Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(Sprintf(":%d", config.Porta), r))
}
