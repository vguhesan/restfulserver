package restfulserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func helloEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: hello")
	message := "{\"hello\"}"
	fmt.Fprintf(w, message)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeEndpoint)
	router.HandleFunc("/hello", helloEndpoint)
	fmt.Println("Server listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}