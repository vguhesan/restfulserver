package restfulserver

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHelloEndPoint(t *testing.T) {
	go handleRequests()

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	r, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, []byte("{\"hello\"}"), body)
}