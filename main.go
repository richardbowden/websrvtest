package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"log"
)

//dont normally like globals, this will never change during the life time of the process
var hostName string

func init() {
	var err error
	hostName, err = os.Hostname()
	if err != nil {
		log.Fatal("cannot get host name of current server")
	}
}

func logRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startReqTime := time.Now()
		f(w, r)
		finReqTime := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), finReqTime.Sub(startReqTime))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I'm served from %s!, this is the dev version", r.Header.Get("X-Forwarded-For"))
}

func main() {
	log.SetOutput(os.Stdout)
	log.Printf("server starting on host %v", hostName)
	http.HandleFunc("/", logRequest(handler))
	log.Fatal(http.ListenAndServe(":8484", nil))
}
