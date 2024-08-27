package rotas

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar Coloca todas as rotas dentro do router.
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotasLogin)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {

			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)

		} else {

			r.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}
