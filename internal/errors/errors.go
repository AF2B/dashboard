package errors

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e APIError) Error() string {
	return e.Message
}

func (e APIError) ToJSON() ([]byte, error) {
	return json.MarshalIndent(e, "", "  ")
}

func (e APIError) RespondWithError(w http.ResponseWriter, err APIError) {
	jsonResponse, _ := err.ToJSON()
	http.Error(w, string(jsonResponse), err.StatusCode)
}

var (
	ErrBadRequest        = APIError{StatusCode: 400, Message: "Requisição inválida"}
	ErrInternalServerError = APIError{StatusCode: 500, Message: "Erro interno do servidor"}
	ErrParsingJSON       = APIError{StatusCode: 500, Message: "Falha ao decodificar o JSON"}
	ErrSavingUser        = APIError{StatusCode: 500, Message: "Falha ao salvar o usuário no banco de dados"}
)