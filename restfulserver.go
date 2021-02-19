package restfulserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// ServerHook has a pointer handle to shutdown the webserver. This enables testing the code modular
type ServerHook struct {
	srv *http.Server
}

var srvHook ServerHook

// Default root handler
func homeEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: home")
}

// Hello handler
func helloEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: hello")
	message := "{\"hello\"}"
	fmt.Fprintf(w, message)
}

// Concatenate a set of strings
type concatenate struct {
	List []string `json:"List"`
}

type catmsg struct {
    Result string `json:"Result"`
}

// Concat handler
func concatEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: concat")
	var cat concatenate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cat); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()
	result := catmsg{ Result : strings.Join(cat.List,"")}
	respondWithJSON(w, http.StatusCreated, result)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeEndpoint)
	router.HandleFunc("/hello", helloEndpoint)
	router.HandleFunc("/concat", concatEndpoint).Methods("POST")
	fmt.Println("Server listening on port 8080")
	srvHook.srv = &http.Server{
        Handler:      router,
        Addr:         "127.0.0.1:8080",
        WriteTimeout: 5 * time.Second,
        ReadTimeout:  3 * time.Second,
    }
	log.Fatal(srvHook.srv.ListenAndServe())
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func init() {
	println("Starting webserver on port:8080 . . .")
	go handleRequests()
}

func main() {
	// doing nothing
}