package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"remoManager/internal/api"
	"remoManager/internal/logger"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	logger.AccessInfoPrint("/", "GET", r.Header.Get("X-Forwarded-For"))
	fmt.Fprintf(w, "It is the root page.\n")
}

func getSensorHandler(w http.ResponseWriter, r *http.Request) {
	logger.AccessInfoPrint("/sensor", "GET", r.Header.Get("X-Forwarded-For"))
	outputStruct := api.GetSensorValues()
	outputJson, err := json.Marshal(outputStruct)
	if err != nil {
		logrus.Errorf("Json Marshal error : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, outputJson)
}

func ServerStart() (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.Methods("GET").Path("/sensor").HandlerFunc(getSensorHandler)
	logger.InfoPrint("API Start")
	http.ListenAndServe(":80", router)
	return err
}
