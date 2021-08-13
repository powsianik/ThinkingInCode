package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type sendData struct {
	key string
	value string
}

var theTests = []struct{
	name string
	url string
	method string
	params []sendData
	expectedStatusCode int
}{
	{"home", "/", "GET", []sendData{}, http.StatusOK},
	{"about", "/about", "GET", []sendData{}, http.StatusOK},
	{"post", "/post", "GET", []sendData{}, http.StatusOK},
	{"posts", "/posts", "GET", []sendData{}, http.StatusOK},
	{"createPost", "/createPost", "GET", []sendData{}, http.StatusOK},
}

func TestHandlers(t *testing.T){
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests{
		if e.method == "GET"{
			response, err := ts.Client().Get(ts.URL + e.url)
			if err != nil{
				t.Log(err)
				t.Fatal(err)
			}

			if response.StatusCode != e.expectedStatusCode{
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, response.StatusCode)
			}
		} else {

		}
	}
}