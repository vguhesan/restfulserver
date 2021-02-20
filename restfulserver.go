package restfulserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/hako/durafmt"
)

// ServerHook has a pointer handle to shutdown the webserver. This enables testing the code modular
type ServerHook struct {
	srv *http.Server
	serverStartTime time.Time
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

type timeRequest struct {
	FutureTime string `json:"FutureTime"`
}

type timeResponse struct {
	Duration string `json:"Duration"`
}

type errorJSONResponse struct {
	Error string `json:"Error"`
}

func subtractTime(time1,time2 time.Time) string { 
	diff := time2.Sub(time1)
	duration := durafmt.Parse(diff.Round(time.Minute)).LimitFirstN(2)
	return fmt.Sprint(duration)
}

// UptimeDiff handler
func futureUptimeEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: futureuptime")
	var timereq timeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&timereq); err != nil {
		var error errorJSONResponse
		error.Error = "Invalid request payloadd"
		respondWithJSON(w, http.StatusCreated, error)
        return
    }
    defer r.Body.Close()
	timeStart := srvHook.serverStartTime
	timeReqInTime, err := time.Parse(time.RFC3339, timereq.FutureTime)
	if err != nil {
		var error errorJSONResponse
		error.Error = "Unparsable Payloadd"
		respondWithJSON(w, http.StatusCreated, error)
        return
	}
	dur := subtractTime(timeStart,timeReqInTime)
	result := timeResponse{ Duration : dur}
	respondWithJSON(w, http.StatusCreated, result)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeEndpoint)
	router.HandleFunc("/hello", helloEndpoint)
	router.HandleFunc("/concat", concatEndpoint).Methods("POST")
	router.HandleFunc("/futureuptime", futureUptimeEndpoint).Methods("POST")
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
	srvHook.serverStartTime = time.Now()
}

func main() {
	// doing nothing
}