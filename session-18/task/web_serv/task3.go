package web_serv

import (
	"fmt"
	"log"
	"net/http"
)

func handlerHomePage(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("ActionLog.Start.HomePage")
	_, err := w.Write([]byte("Welcome to the homepage!"))
	if err != nil {
		log.Default().Println(err.Error())
	}
	log.Default().Println("ActionLog.End.HomePage")
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("ActionLog.Start.About")
	_, err := w.Write([]byte("This is the about page."))
	if err != nil {
		log.Default().Println(err.Error())
	}
	log.Default().Println("ActionLog.End.About")
}

func handlerContact(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("ActionLog.Start.Contact")
	_, err := w.Write([]byte("Contact us at contact@example.com."))
	if err != nil {
		log.Default().Println(err.Error())
	}
	log.Default().Println("ActionLog.End.Contact")
}
func Task3() {
	fmt.Println("Task3   *****************")

	http.HandleFunc("/", handlerHomePage)
	http.HandleFunc("/about", handlerAbout)
	http.HandleFunc("/contact", handlerContact)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
