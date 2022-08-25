package main

import (
	"encoding/json"
	"fileFinderServer/fileFinder"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Incoming JSON struct for `largeFilesHandler`
type filesHandlerJSON struct {
	Path  string `json:"path"`
	Limit int    `json:"limit"`
}

func largeFilesHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var reqBody filesHandlerJSON
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Endcoding, Content-Type, Content-Length, Authorization, X-CSRF-token")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")

	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err.Error() == "EOF" {
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Request body did not provide the necessary paramaeters"})
			return
		}
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	files, err := fileFinder.GetTopXFiles(reqBody.Path, reqBody.Limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(files)

}

func deleteFileHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Endcoding, Content-Type, Content-Length, Authorization, X-CSRF-token")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, DELETE")
	if r.Method == http.MethodOptions {
		return
	}

	var reqBody fileFinder.FileDisplay
	if r.Method != http.MethodDelete {
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "This endpoint only accepts delete requests"})
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err.Error() == "EOF" {
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Request body did not provide the necessary paramaeters"})
			return
		}
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	err = reqBody.Delete()
	if err != nil {
		logrus.Error(reqBody.Path)

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(reqBody)

}
