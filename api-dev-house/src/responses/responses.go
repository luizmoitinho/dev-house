package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON ... retorna uma resposta em JSON para a request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

//Error ... retorna um erro em formato JSON
func Error(w http.ResponseWriter, statusCode int, err error) {

	JSON(w, statusCode, struct {
		Erro string `json:"erro,omitempty"`
	}{
		Erro: err.Error(),
	})

}
