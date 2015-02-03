package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
)

func HomeHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	w.Header().Set("X-Tropo-Backend-Path", req.URL.Path)
	w.Header().Set("X-Tropo-Backend-Protocol", req.Proto)
	w.Header().Set("X-Tropo-Backend", req.Header.Get("X-Tropo-Backend"))

	var rest = regexp.MustCompile(`^/rest.*`)
	var session = regexp.MustCompile(`^/session.*`)

	switch {
	case rest.MatchString(req.URL.Path):
		log.Debug("Path: " + req.URL.Path + " matches provisioning")
		fmt.Fprintf(w, "you got rest\n")
	case session.MatchString(req.URL.Path):
		log.Debug("Path: " + req.URL.Path + " matches sessions")
		fmt.Fprintf(w, "you got sessions\n")
	default:
		log.Debug("Path: " + req.URL.Path + " has no matches")
		http.Error(w, "fail boat\n", http.StatusServiceUnavailable)
	}

}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	level, err := log.ParseLevel("debug")
	if err != nil {
		log.Fatal(err)
	}
	applicationLogLevel := level.String()
	log.Warn("Logging at " + applicationLogLevel + " level")
	log.SetLevel(level)
}

func main() {
	router := httprouter.New()
	router.GET("/*wildcard", HomeHandler)
	log.Fatal(http.ListenAndServe(":8081", router))
}
