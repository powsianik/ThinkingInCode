package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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
	{"createPost", "/createPost", "POST", []sendData{
		{key: "creatorName", value: "Test"},
		{key: "title", value: "Test"},
		{key: "description", value: "Test"},
		{key: "content", value: "Test"},
		{key: "image", value: "Test"},
	}, http.StatusOK},
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
			values := url.Values{}
			for _, x := range e.params{
				values.Add(x.key, x.value)
			}

			response, err := ts.Client().PostForm(ts.URL + e.url, values)
			if err != nil{
				t.Log(err)
				t.Fatal(err)
			}

			if response.StatusCode != e.expectedStatusCode{
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, response.StatusCode)
			}
		}
	}
}