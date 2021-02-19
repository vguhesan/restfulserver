package restfulserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ServerHook has a pointer handle to shutdown the webserver. This enables testing the code modular
type ServerHook struct {
	srv *http.Server
}

var srvHook ServerHook

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
	srvHook.srv = &http.Server{
        Handler:      router,
        Addr:         "127.0.0.1:8080",
        WriteTimeout: 5 * time.Second,
        ReadTimeout:  3 * time.Second,
    }
	log.Fatal(srvHook.srv.ListenAndServe())
}

func init() {
	println("Starting webserver on port:8080 . . .")
	go handleRequests()
}

func main() {
	// doing nothing
}