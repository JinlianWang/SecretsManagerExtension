package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// SecretResponse is the response struct for the secret value
type SecretResponse struct {
	SecretValue string `json:"secretValue"`
}

// Handler handles the Lambda function invocation
func Handler(ctx context.Context) (SecretResponse, error) {
	// Initialize the AWS session
	sess, err := session.NewSession()
	if err != nil {
		return SecretResponse{}, err
	}

	// Initialize the Secrets Manager client
	svc := secretsmanager.New(sess)

	// Get the secret value from AWS Secrets Manager
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(os.Getenv("SECRET_ID")),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return SecretResponse{}, err
	}

	// Construct the secret response
	secretResponse := SecretResponse{
		SecretValue: *result.SecretString,
	}

	// Return the secret response
	return secretResponse, nil
}

func main() {
	// Set up the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := Handler(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the secret response as JSON
		json.NewEncoder(w).Encode(response)
	})

	// Start the HTTP server
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
