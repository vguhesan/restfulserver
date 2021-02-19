package restfulserver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var srvURL = "http://localhost:8080"

func TestHelloEndPoint(t *testing.T) {
	
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	r, _ := http.NewRequest("GET", srvURL+"/hello", nil)
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

func TestConcatEndPoint(t *testing.T) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	var jsonStr = []byte(
		`
		{
			"List": [
			"abcd",
			"egh",
			"ijkmnop"
			]
		}	
		`)
	r, _ := http.NewRequest("POST", srvURL+"/concat", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}
	json.Unmarshal(body, &m)
	var expectedValue = "abcdeghijkmnop"
	var actualValue = m["Result"]
	assert.Equal(t, expectedValue, actualValue)
}

