package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// API handler
func apiHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, world!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api", apiHandler)

	port := 8080
	fmt.Printf("API is running on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
