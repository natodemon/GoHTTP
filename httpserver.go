package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	portNo := flag.Int("p", 8080, "port to listen for requests")

	flag.Parse()

	router := http.NewServeMux()

	htServer := http.FileServer(http.Dir("./web"))
	router.Handle("/", htServer)

	println("HTTP server started on port", *portNo)

	var pString string = ":" + strconv.Itoa(*portNo)
	log.Fatal(http.ListenAndServe(pString, requestLogger(router)))
}

func requestLogger(hand http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqTime := time.Now()

		hand.ServeHTTP(w, r)

		fmt.Printf("%v\t\t%s\t%s\t%s\t\n", reqTime.Format(time.Stamp), r.RequestURI, r.Method, r.RemoteAddr)
	})
}
