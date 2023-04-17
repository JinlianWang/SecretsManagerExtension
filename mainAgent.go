package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// SecretResponse is the response struct for the secret value
type SecretResponse struct {
	SecretValue string `json:"secretValue"`
}

// Handler handles the HTTP request
func Handler(w http.ResponseWriter, r *http.Request) {
	// Get the secret ID from the URL path
	segments := strings.Split(r.URL.Path, "/")
	secretID := strings.Join(segments[2:], "/") // join every path component after the /secrets/ prefix
	if secretID == "" {
		http.Error(w, "Missing secretID parameter", http.StatusBadRequest)
		return
	}

	// Print the secret ID for debugging
	fmt.Println("Secret ID:", secretID)

	// Initialize the AWS session
	sess, err := session.NewSession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Initialize the Secrets Manager client
	svc := secretsmanager.New(sess)

	// Get the secret value from AWS Secrets Manager
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Construct the secret response
	secretResponse := SecretResponse{
		SecretValue: *result.SecretString,
	}

	// Write the secret response as JSON
	json.NewEncoder(w).Encode(secretResponse)
}

func main() {
	// Set up the HTTP server
	http.HandleFunc("/secrets/", Handler)

	// Start the HTTP server
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
