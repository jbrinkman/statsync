package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/valkey-io/valkey-go"
)

func main() {
	addr := os.Getenv("VALKEY_ADDR")
	if addr == "" {
		addr = "valkey:6379"
	}

	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{addr}})
	if err != nil {
		log.Fatalf("failed to connect to valkey: %v", err)
	}
	defer client.Close()

	// Health check
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		res := client.Do(context.Background(), client.B().Ping().Build())
		if res.Error() != nil {
			http.Error(w, "valkey unavailable", http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Println("backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
