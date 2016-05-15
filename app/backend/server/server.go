package server

import (
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func handlePrefix(mux *mux.Router, url string, handler http.Handler ) {
	mux.PathPrefix(url).Handler(http.StripPrefix(url, handler))
}

func handleStaticServer (mux *mux.Router) {
	handlePrefix(mux, "/static/", http.FileServer(http.Dir("./frontend/build/web/static")))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile("./frontend/build/web/index.html")
	if (err != nil) {
		w.Write([]byte("Erro!"))
	} else {
		w.Write(bytes)
	}
}

func BackendHandler() http.Handler {
	mux := mux.NewRouter()
	handleStaticServer(mux)
	handlePrefix(mux, "/", http.HandlerFunc(handleIndex))
	return mux
}

