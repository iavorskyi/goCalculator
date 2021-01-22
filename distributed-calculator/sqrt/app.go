// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

type Operands struct {
	OperandOne float32 `json:"operandOne,string"`
	OperandTwo float32 `json:"operandTwo,string"`
}

func sqrt(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in sqrt main")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var operands Operands
	err := json.NewDecoder(r.Body).Decode(&operands)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Println(fmt.Sprintf("%s%f", "Getting sqroot ", operands.OperandOne))
	err = json.NewEncoder(w).Encode(math.Sqrt(float64(operands.OperandOne)))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/sqrt", sqrt).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":6000", router))
}
