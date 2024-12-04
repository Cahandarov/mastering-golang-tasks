package intro_net_http

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("ActionLog.Start")
	w.Header().Set("X-Custom-Header", "Learning Go")
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Println("Error happened on response: ", err)
	}

	log.Default().Println("ActionLog.End")

}

func Task1() {
	fmt.Println("Task1   *****************")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
