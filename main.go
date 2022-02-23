package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("vim-go")
	r := mux.NewRouter()

	// - Get
	getR := r.PathPrefix("/get").
		Methods("GET").
		Subrouter()

	getR.HandleFunc("", NoBodyHandler)
	getR.HandleFunc("/{pthVar0}", PathVarHandler)
	getR.HandleFunc("/{pthVar0}/var2/{pthVar1}", TwoPathVarHandler)
	getR.HandleFunc("/{pthVar0}/var2/{pthVar1}/closing", TwoPathVarHandler)

	// - Post
	pstR := r.PathPrefix("/post").
		Methods("POST").
		Subrouter()

	pstR.HandleFunc("", NoBodyHandler)
	pstR.HandleFunc("/{pthVar0}", PathVarHandler)
	pstR.HandleFunc("/{pthVar0}/var2/{pthVar1}", TwoPathVarHandler)
	pstR.HandleFunc("/{pthVar0}/var2/{pthVar1}/closing", TwoPathVarHandler)

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

type SampleRequest struct {
	value string
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
		res.Write([]byte("Successful one param"))
		return
	}
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("Unsuccessful one param"))
}

func TwoPathVarHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	if vars["pthVar0"] == "valid" && vars["pthVar1"] == "valid" {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Successful two param"))
		return
	}
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("Unsuccessful two param"))
}

func BodyHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	var b SampleRequest

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		res.WriteHeader(http.StatusUnprocessableEntity)
		res.Write([]byte("Failed to parse"))
		return
	}
	if b.value != "valid" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid request"))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Successful body no params"))

}
