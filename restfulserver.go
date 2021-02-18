package main

import (
	"fmt"
	"log"
	"net/http"
)


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func helloEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: hello")
	message := "{\"hello\"}"
	fmt.Fprintf(w, message)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", helloEndpoint)
    log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server listening on port 8080")
}

func main() {
	handleRequests()
}