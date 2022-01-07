package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//exit process immediately upon sigterm
	handleSigTerms()

	//start http server
	startServer()
}

func startServer() {
	port := getenv("PORT", "8080")
	fmt.Printf("PORT = %v \n", port)

	hc := getenv("HEALTHCHECK", "/health")
	fmt.Printf("HEALTHCHECK = %v \n", hc)

	log.Printf("http server listening on %s \n", port)
	http.HandleFunc(hc, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusNotImplemented)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func handleSigTerms() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(1)
	}()
}
