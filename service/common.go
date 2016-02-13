package service

import (
	"encoding/json"
	"log"
	"net/http"
)

// StandardErrorRespModel ...
type StandardErrorRespModel struct {
	ErrorMessage string `json:"error"`
}

// -----------------
// --- Successes ---

// RespondWithSuccessOK ...
func RespondWithSuccessOK(w http.ResponseWriter, respModel interface{}) {
	RespondWithSuccess(w, http.StatusOK, respModel)
}

// RespondWithSuccess ...
func RespondWithSuccess(w http.ResponseWriter, httpStatusCode int, respModel interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	if err := json.NewEncoder(w).Encode(&respModel); err != nil {
		log.Println(" [!] Exception: respondWithSuccess: Error: ", err)
	}
}

// --------------
// --- Errors ---

// RespondWithBadRequestError ...
func RespondWithBadRequestError(w http.ResponseWriter, errMsg string) {
	RespondWithError(w, http.StatusBadRequest, errMsg)
}

// RespondWithNotFoundError ...
func RespondWithNotFoundError(w http.ResponseWriter, errMsg string) {
	RespondWithError(w, http.StatusNotFound, errMsg)
}

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, httpErrCode int, errMsg string) {
	resp := StandardErrorRespModel{
		ErrorMessage: errMsg,
	}
	RespondWithErrorJSON(w, httpErrCode, resp)
}

// RespondWithErrorJSON ...
func RespondWithErrorJSON(w http.ResponseWriter, httpErrCode int, respModel interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErrCode)
	if err := json.NewEncoder(w).Encode(&respModel); err != nil {
		log.Println("Error: ", err)
	}
}
