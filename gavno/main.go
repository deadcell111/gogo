package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Body struct {
	mass float64
	velocity float64
}

func (r *Body) momentum() float64 {
	return(r.mass * r.velocity)
}

func example (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mass, err := strconv.ParseFloat(vars["mass"], 64)
	if err != nil {
		panic("Nuh-uh...01")
	}
	velocity, err := strconv.ParseFloat(vars["velocity"], 64)
	if err != nil {
		panic("Nuh-uh...02")
	}
	bot := Body{mass, velocity}
	jsonResponse, jsonError := json.Marshal(bot.momentum())
	if jsonError != nil {
		panic("Nuh-uh...03")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/example/{mass:[0-9]+}/{velocity:[0-9]+}/", example)
	log.Fatal(http.ListenAndServe(":8080", r))
}