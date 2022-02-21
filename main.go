package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("vim-go")
	r := mux.NewRouter()
	getR := r.PathPrefix("/get").
		Methods("GET").
		Subrouter()

	getR.HandleFunc("/", GetHandler)
	getR.HandleFunc("/{pthVar0}", GetPathVarHandler)
	getR.HandleFunc("/{pthVar0}/var2/{pthVar1}", Get2PathVarHandler)
	getR.HandleFunc("/{pthVar0}/var2/{pthVar1}/closing", Get2PathVarHandler)

	pstR := r.PathPrefix("/post").Subrouter()
	pchR := r.PathPrefix("/patch").Subrouter()
	putR := r.PathPrefix("/put").Subrouter()
	delR := r.PathPrefix("/delete").Subrouter()

}

func GetHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	res.Write([]byte("Successful GET"))
}

func GetPathVarHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	if vars["pthVar0"] == "valid" {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Successful GET"))
	}
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("Unsuccessful GET"))
}

func Get2PathVarHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	if vars["pthVar0"] == "valid" && vars["pthVar1"] == "valid" {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Successful GET"))
	}
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("Unsuccessful GET"))
}
