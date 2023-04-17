package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/event/next", handleNextEvent)
	http.HandleFunc("/ext-server", handleExtServer)

	go func() {
		log.Println("Starting extension server")
		if err := http.ListenAndServe(":"+os.Getenv("AWS_LAMBDA_RUNTIME_API"), nil); err != nil {
			log.Fatalf("The HTTP request failed with error %s\n", err)
		}
	}()

	log.Println("Starting external HTTP server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	}
}

func handleNextEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling event")

	// Read and log the eventType
	eventType := r.Header.Get("Lambda-Extension-Function-Error-Type")
	log.Printf("Event type: %s\n", eventType)

	switch eventType {
	case "INVOKE":
		log.Println("Function is being invoked")
	case "SHUTDOWN":
		log.Println("Function is shutting down")
		os.Exit(0)
	}

	w.WriteHeader(http.StatusOK)
}

func handleExtServer(w http.ResponseWriter, r *http.Request) {
	message := "Hello from Lambda extension HTTP server!"
	fmt.Fprint(w, message)
}
