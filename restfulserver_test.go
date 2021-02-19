package restfulserver

import (
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

