package web_serv

import (
	"fmt"
	"log"
	"net/http"
)

func handleLog(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("ActionLog.StartLogging")
	method := r.Method
	url := r.URL
	fmt.Printf("Method: %s, URL: %s\n", method, url)
	_, err := w.Write([]byte("Request logged."))
	if err != nil {
		log.Default().Println(err.Error())
	}
	log.Default().Println("ActionLog.EndLogging")
}

func Task4() {
	fmt.Println("Task4   *****************")
	http.HandleFunc("/log", handleLog)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
