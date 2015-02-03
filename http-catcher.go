package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func BackendHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("X-Tropo-Backend-Path", req.URL.Path)
	w.Header().Set("X-Tropo-Backend-Protocol", req.Proto)
	w.Header().Set("X-Tropo-Backend", req.Header.Get("X-Tropo-Backend"))
	fmt.Fprintf(w, "you got %s\n", ps.ByName("name"))
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

	client := httprouter.New()
	client.GET("/*name", BackendHandler)

	http.ListenAndServe("localhost:9082", client)

}
