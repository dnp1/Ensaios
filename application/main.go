package main

import (
	"github.com/daniloanp/Ensaios/application/backend/server"
	"net/http"
	"flag"
	"github.com/daniloanp/Ensaios/application/backend/bootstrap"
)


var (

	flags struct{bootstrap bool}

)
func main() {
	flag.BoolVar(&flags.bootstrap, "do-bootstrap", false, "")
	flag.Parse()
	if flags.bootstrap {
		bootstrap.Bootstrap()
	} else {
		mux := server.BackendHandler()
		http.ListenAndServe(":8080", mux)
	}

}