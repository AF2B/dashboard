package handlers

import (
	"dashboard/internal/models"
	"dashboard/internal/repository"
	"encoding/json"
	"net/http"
	"time"
)

func ActivityVisitors(w http.ResponseWriter, r *http.Request) {
	var visitor models.ActivityVisitor
	err := json.NewDecoder(r.Body).Decode(&visitor)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}

	visitor.Date = time.Now()

	err = repository.CreateActivityVisitor(&visitor)
	if err != nil {
		http.Error(w, "Erro ao criar visitante", http.StatusInternalServerError)
        return
	}

	 w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(visitor)
}