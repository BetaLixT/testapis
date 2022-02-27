package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	getR.HandleFunc("/oq", NoBodyOneQeryHandler)
	getR.HandleFunc("/tq", NoBodyTwoQeryHandler)
	getR.HandleFunc("/oq/{pthVar0}", NoBodyOneQeryOneParamHandler)
	getR.HandleFunc("/tq/{pthVar0}", NoBodyOneQeryOneParamHandler)
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

	r.HandleFunc("/", NoBodyHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Listening to port %s...\n", port)
	log.Fatal(srv.ListenAndServe())
}

type SampleRequest struct {
	Value string `json:"value"`
}

type SampleResponse struct {
	Response string `json:"response"`
	Success  bool   `json:"success"`
}

// TODO these functions can be a lot less stupid
// but this was just a dumb throw away app to learn
// stuff soo.... TODON'T?
func NoBodyHandler(
	res http.ResponseWriter,
	req *http.Request,
) {

	writeResponse(
		&res,
		http.StatusOK,
		"Successful No body",
		true,
	)
}

func NoBodyOneQeryHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	query := req.URL.Query()
	var0 := query.Get("var0")

	if var0 == "valid" {
		writeResponse(
			&res,
			http.StatusOK,
			"Successful No body one query",
			true,
		)
		return
	}

	writeResponse(
		&res,
		http.StatusNotFound,
		"Unsuccessful No body one query",
		false,
	)
}

func NoBodyTwoQeryHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	query := req.URL.Query()
	var0 := query.Get("var0")
	var1 := query.Get("var1")

	if var0 == "valid" && var1 == "valid" {
		writeResponse(
			&res,
			http.StatusOK,
			"Successful No body two query",
			true,
		)
		return
	}

	writeResponse(
		&res,
		http.StatusNotFound,
		"Unsuccessful No body two query",
		false,
	)
}

func PathVarHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	if vars["pthVar0"] == "valid" {
		writeResponse(
			&res,
			http.StatusOK,
			"Successful one param",
			true,
		)
		return
	}

	writeResponse(
		&res,
		http.StatusNotFound,
		"Unsuccessful one param",
		false,
	)
}

func NoBodyOneQeryOneParamHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	query := req.URL.Query()
	var0 := query.Get("var0")

	if var0 != "valid" {
		writeResponse(
			&res,
			http.StatusNotFound,
			"Unsuccessful No body one query",
			false,
		)
		return
	}
	vars := mux.Vars(req)
	if vars["pthVar0"] != "valid" {
		writeResponse(
			&res,
			http.StatusNotFound,
			"Unsuccessful No body one query",
			false,
		)
		return
	}

	writeResponse(
		&res,
		http.StatusOK,
		"Successful No body one query",
		true,
	)
}

func NoBodyTwoQeryOneParamHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	query := req.URL.Query()
	var0 := query.Get("var0")
	var1 := query.Get("var1")

	if var0 != "valid" || var1 != "valid" {
		writeResponse(
			&res,
			http.StatusNotFound,
			"Unsuccessful No body two query",
			false,
		)
		return
	}

	vars := mux.Vars(req)
	if vars["pthVar0"] != "valid" {
		writeResponse(
			&res,
			http.StatusNotFound,
			"Unsuccessful No body one query",
			false,
		)
		return
	}
	writeResponse(
		&res,
		http.StatusOK,
		"Successful No body two query",
		true,
	)
}

func TwoPathVarHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	vars := mux.Vars(req)
	if vars["pthVar0"] == "valid" && vars["pthVar1"] == "valid" {
		writeResponse(
			&res,
			http.StatusOK,
			"Successful two param",
			true,
		)
		return
	}

	writeResponse(
		&res,
		http.StatusNotFound,
		"Unsuccessful two param",
		false,
	)
}

func BodyHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	var b SampleRequest

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		writeResponse(
			&res,
			http.StatusUnprocessableEntity,
			"Failed to parse",
			false,
		)
		return
	}

	if b.Value != "valid" {
		writeResponse(
			&res,
			http.StatusBadRequest,
			fmt.Sprintf("Invalid request %s", b.Value),
			false,
		)
		return
	}

	writeResponse(
		&res,
		http.StatusOK,
		"Successful body no params",
		true,
	)
}

func BodyOneParamHandler(
	res http.ResponseWriter,
	req *http.Request,
) {
	var b SampleRequest

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		writeResponse(
			&res,
			http.StatusUnprocessableEntity,
			"Failed to parse",
			false,
		)
		return
	}
	vars := mux.Vars(req)
	param1 := vars["pthVar0"]
	if b.Value != "valid" {

		writeResponse(
			&res,
			http.StatusBadRequest,
			"Invalid request",
			false,
		)
		return
	}

	if param1 != "valid" {

		writeResponse(
			&res,
			http.StatusNotFound,
			"Unsuccesful body one param",
			false,
		)
		return
	}

	writeResponse(
		&res,
		http.StatusOK,
		"Successful body one param",
		true,
	)

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

	if b.Value != "valid" {
		writeResponse(
			&res,
			http.StatusBadRequest,
			"Invalid request",
			false,
		)
		return
	}

	if param1 != "valid" || param2 != "valid" {

		writeResponse(
			&res,
			http.StatusNotFound,
			"Unsuccesful body two params",
			false,
		)
		return
	}

	writeResponse(
		&res,
		http.StatusOK,
		"Successful body two params",
		true,
	)
}

func writeResponse(
	res *http.ResponseWriter,
	statusCode int,
	response string,
	success bool,
) {

	pld, err := json.Marshal(SampleResponse{
		Response: response,
		Success:  success,
	})
	if err != nil {
		// Should not be happening...
		(*res).WriteHeader(http.StatusInternalServerError)
		(*res).Write([]byte("Failed to marshal response"))
		return
	}
	(*res).Header().Set("Content-Type", "application/json")
	(*res).WriteHeader(statusCode)
	(*res).Write(pld)
}
