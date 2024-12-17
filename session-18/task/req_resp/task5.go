package req_resp

import (
	"fmt"
	"log"
	"net/http"
)

func handleGreet(w http.ResponseWriter, req *http.Request) {
	log.Default().Println("ActionLog.StartGreeting")
	name := req.URL.Query().Get("name")
	if name == "" {
		_, err := w.Write([]byte("Hello, Stranger!"))
		if err != nil {
			log.Default().Println(err.Error())
		}
	} else {
		_, errN := w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
		if errN != nil {
			log.Default().Println(errN.Error())
		}
	}

	log.Default().Println("ActionLog.EndGreeting")
}

func Task5() {
	fmt.Println("Task5   *****************")

	http.HandleFunc("/greet", handleGreet)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
