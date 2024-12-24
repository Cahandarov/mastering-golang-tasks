package util

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"session-22/model"
)

func handleError(w http.ResponseWriter, err *model.ErrorResponse) {
	w.WriteHeader(err.Code)
	encodeErr := json.NewEncoder(w).Encode(err)
	if encodeErr != nil {
		log.Error("ActionLog.HandleError.error while encoding JSON:", encodeErr)
	}
}

func PrepareResponse(w http.ResponseWriter, data interface{}, err *model.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(err.Code) // Correctly set the HTTP status code
		if encodeErr := json.NewEncoder(w).Encode(err); encodeErr != nil {
			log.Error("Failed to encode error response:", encodeErr)
		}
		return
	}

	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	encodeErr := json.NewEncoder(w).Encode(data)
	if encodeErr != nil {
		log.Error("ActionLog.PrepareResponse.error while encoding JSON:", encodeErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
