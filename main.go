package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Strating server on http://localhost:8000")
	router := mux.NewRouter()
	router.HandleFunc("/large-files", largeFilesHandler).Methods("OPTIONS", "POST")
	router.HandleFunc("/delete-file", deleteFileHandler).Methods("DELETE", "OPTIONS")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	logrus.SetLevel(logrus.ErrorLevel)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		logrus.Fatal(err)
	}

}
