package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/*func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}*/

func isFileEmpty(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err // file does not exist or other error
	}
	return info.Size() == 0, nil
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	empty, err := isFileEmpty("/tmp/ready.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if empty {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Readiness FAILED")
		fmt.Fprintln(w, "Readiness FAILED")
	} else {
		w.WriteHeader(http.StatusOK)
		log.Printf("Readiness OK")
		fmt.Fprintln(w, "Readiness OK")
	}
}

func livezHandler(w http.ResponseWriter, r *http.Request) {
	empty, err := isFileEmpty("/tmp/live.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if empty {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Liveness FAILED")
		fmt.Fprintln(w, "Liveness FAILED")
	} else {
		w.WriteHeader(http.StatusOK)
		log.Printf("Liveness OK")
		fmt.Fprintln(w, "Liveness OK")

	}
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/livez", livezHandler)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
