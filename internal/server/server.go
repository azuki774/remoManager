package server

import (
	"fmt"
	"net/http"
	"remoManager/internal/logger"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	logger.AccessInfoPrint("/mawinter/", "GET", r.Header.Get("X-Forwarded-For"))
	fmt.Fprintf(w, "It is the root page.\n")
}

func ServerStart() (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	logger.InfoPrint("API Start")
	http.ListenAndServe(":80", router)
	return err
}
