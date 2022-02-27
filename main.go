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

	pstR.HandleFunc("/body", BodyHandler)
	pstR.HandleFunc("/body/{pthVar0}", BodyOneParamHandler)
	pstR.HandleFunc("/body/{pthVar0}/var2/{pthVar1}", BodyTwoParamHandler)
	pstR.HandleFunc("/body/{pthVar0}/var2/{pthVar1}/closing", BodyTwoParamHandler)

	pstR.HandleFunc("", NoBodyHandler)
	pstR.HandleFunc("/{pthVar0}", PathVarHandler)
	pstR.HandleFunc("/{pthVar0}/var2/{pthVar1}", TwoPathVarHandler)
	pstR.HandleFunc("/{pthVar0}/var2/{pthVar1}/closing", TwoPathVarHandler)

	// - Put
	putR := r.PathPrefix("/put").
		Methods("PUT").
		Subrouter()

	putR.HandleFunc("/body", BodyHandler)
	putR.HandleFunc("/body/{pthVar0}", BodyOneParamHandler)
	putR.HandleFunc("/body/{pthVar0}/var2/{pthVar1}", BodyTwoParamHandler)
	putR.HandleFunc("/body/{pthVar0}/var2/{pthVar1}/closing", BodyTwoParamHandler)

	putR.HandleFunc("", NoBodyHandler)
	putR.HandleFunc("/{pthVar0}", PathVarHandler)
	putR.HandleFunc("/{pthVar0}/var2/{pthVar1}", TwoPathVarHandler)
	putR.HandleFunc("/{pthVar0}/var2/{pthVar1}/closing", TwoPathVarHandler)

	// - Patch
	pchR := r.PathPrefix("/patch").
		Methods("PATCH").
		Subrouter()

	pchR.HandleFunc("/body", BodyHandler)
	pchR.HandleFunc("/body/{pthVar0}", BodyOneParamHandler)
	pchR.HandleFunc("/body/{pthVar0}/var2/{pthVar1}", BodyTwoParamHandler)
	pchR.HandleFunc("/body/{pthVar0}/var2/{pthVar1}/closing", BodyTwoParamHandler)

	pchR.HandleFunc("", NoBodyHandler)
	pchR.HandleFunc("/{pthVar0}", PathVarHandler)
	pchR.HandleFunc("/{pthVar0}/var2/{pthVar1}", TwoPathVarHandler)
	pchR.HandleFunc("/{pthVar0}/var2/{pthVar1}/closing", TwoPathVarHandler)

	// - Delete
	delR := r.PathPrefix("/delete").
		Methods("DELETE").
		Subrouter()

	delR.HandleFunc("/body", BodyHandler)
	delR.HandleFunc("/body/{pthVar0}", BodyOneParamHandler)
	delR.HandleFunc("/body/{pthVar0}/var2/{pthVar1}", BodyTwoParamHandler)
	delR.HandleFunc("/body/{pthVar0}/var2/{pthVar1}/closing", BodyTwoParamHandler)

	delR.HandleFunc("", NoBodyHandler)
	delR.HandleFunc("/{pthVar0}", PathVarHandler)
	delR.HandleFunc("/{pthVar0}/var2/{pthVar1}", TwoPathVarHandler)
	delR.HandleFunc("/{pthVar0}/var2/{pthVar1}/closing", TwoPathVarHandler)

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
	Value string `json:"value"`
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

	if b.Value != "valid" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Invalid request %s", b.Value)))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Successful body no params"))

}

func BodyOneParamHandler(
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
	vars := mux.Vars(req)
	param1 := vars["pthVar0"]
	if b.Value != "valid" || param1 != "valid" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid request"))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Successful body one param"))

}

func BodyTwoParamHandler(
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
	vars := mux.Vars(req)
	param1 := vars["pthVar0"]
	param2 := vars["pthVar1"]
	isValid := b.Value != "valid" ||
		param1 != "valid" ||
		param2 != "valid"
	if isValid {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid request"))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Successful body two params"))
}
