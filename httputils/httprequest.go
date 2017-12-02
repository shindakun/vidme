package httputils

import (
	// "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Generic HTTP Request Method, returns HTTP status code.
func HttpRequest(method, uri string, cookie string, data io.Reader) []byte {
	client := &http.Client{}
	var send io.Reader

	if data != nil {
		send = data
	} else {
		send = nil
	}

	req, err := http.NewRequest(method, uri, send)
	if err != nil {
		log.Fatal(err)
	}

	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s", body)

	return body
}
