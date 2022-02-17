package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	runPath := "/ko-app/zombies"
	mode := getenv("MODE", "http")
	log.SetPrefix(fmt.Sprintf("%d: ", os.Getpid()))

	http.HandleFunc("/launch", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command(runPath)
		cmd.Env = []string{"MODE=sleep"}
		cmd.Start()
	})
	switch mode {
	case "http":
		log.Fatalf("err serving http: %v", http.ListenAndServe(":8080", nil))
	default:
		time.Sleep(time.Hour)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
