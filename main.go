package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// Define a struct to match the webhook payload structure
type WebhookPayload struct {
	Repository struct {
		Status string `json:"status"`
	} `json:"repository"`
}

func updateContainer() {
	scriptPath := "/update_container.sh"
	cmd := exec.Command("/bin/sh", scriptPath)
	cmd.Env = os.Environ() // Use the current environment variables
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Executing update script...")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing update script: %v\n", err)
	} else {
		fmt.Println("Update script executed successfully.")
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Webhook received!")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Repository status: %s\n", payload.Repository.Status)

	// Assuming "Active" status indicates the final update
	if payload.Repository.Status == "Active" {
		updateContainer()
	}

	fmt.Fprintf(w, "Webhook processed")
}

func main() {
	updateContainer()

	http.HandleFunc("/webhook", handleWebhook)
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
