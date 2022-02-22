package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("vim-go")
	r := mux.NewRouter()
	getR := r.PathPrefix("/get").
		Methods("GET").
		Subrouter()

	getR.HandleFunc("", NoBodyHandler)
	getR.HandleFunc("/{pthVar0}", PathVarHandler)
	getR.HandleFunc("/{pthVar0}/var2/{pthVar1}", TwoPathVarHandler)
	getR.HandleFunc("/{pthVar0}/var2/{pthVar1}/closing", TwoPathVarHandler)

	// pstR := r.PathPrefix("/post").Subrouter()
	// pchR := r.PathPrefix("/patch").Subrouter()
	// putR := r.PathPrefix("/put").Subrouter()
	// delR := r.PathPrefix("/delete").Subrouter()

	r.HandleFunc("/", NoBodyHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8084",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func NoBodyHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	res.Write([]byte("Successful No Body"))
}

func PathVarHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	if vars["pthVar0"] == "valid" {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Successful GET"))
		return
	}
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("Unsuccessful GET"))
}

func TwoPathVarHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	if vars["pthVar0"] == "valid" && vars["pthVar1"] == "valid" {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Successful GET"))
		return
	}
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("Unsuccessful GET"))
}
