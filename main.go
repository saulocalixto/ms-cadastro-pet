package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/saulocalixto/ms-cadastro-pet/internal/dto"
	"github.com/saulocalixto/ms-cadastro-pet/internal/services/async"
	"github.com/saulocalixto/ms-cadastro-pet/internal/services/sync"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc(
		"/cadastrar-async",
		func(w http.ResponseWriter, r *http.Request) { CadastrarPet(w, r,
			func(dtos []dto.PetDTO) []dto.PetValidatedDTO { return async.Salvar(dtos) })}).Methods("POST")
	router.HandleFunc(
		"/cadastrar-sync",
		func(w http.ResponseWriter, r *http.Request) { CadastrarPet(w, r,
			func(dtos []dto.PetDTO) []dto.PetValidatedDTO { return sync.Salvar(dtos) })}).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CadastrarPet(w http.ResponseWriter, r *http.Request, salvar func([]dto.PetDTO) []dto.PetValidatedDTO) {
	var contrato dto.ContratoDTO
	err := json.NewDecoder(r.Body).Decode(&contrato)
	if err != nil {
		log.Panic("Erro ao desserializar o contrato")
	}
	inicio := time.Now()
	validations := salvar(contrato.Mensagens)
	fim := time.Now()
	w.Header().Set("Content-Type", "application/json")
	result := dto.ReturnDTO{Validations: validations, Tempo: fmt.Sprint(fim.Sub(inicio)), Total: len(contrato.Mensagens)}
	payload, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(payload)
}
