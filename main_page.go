package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PostData struct {
	Url, Method string
	Headers     map[string]string
	Data        string
}

func makeRequest(d PostData) http.Request {
	method := strings.ToLower(d.Method)
	switch method {
	case http.MethodGet | http.MethodPost:
		continue
	default:
		log.Printf("invalid method %v\n", method)
	}
	headers := make(map[string][]string, len(d.Headers))
	for k, v := range d.Headers {
		headers[k] = []string{v}
	}
	return http.Request{
		Method:     method,
		RequestURI: d.Url,
		Header:     headers,
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var d PostData
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			log.Println(err)
			return
		}
		req := makeRequest(d)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Errorln(err)
		}

	} else if r.Method == http.MethodGet {

	}
	fmt.Fprintln(w, "ok.")
}
