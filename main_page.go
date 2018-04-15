package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func makeRequest(d PostData) *http.Request {
	method := strings.ToUpper(d.Method)
	switch method {
	case http.MethodGet, http.MethodPost:
		log.Println("making a request...")
	default:
		log.Printf("invalid method %v\n", method)
	}
	headers := make(map[string][]string, len(d.Headers))
	for k, v := range d.Headers {
		headers[k] = []string{v}
	}
	return &http.Request{
		Method:     method,
		RequestURI: d.Url,
		Header:     headers,
	}
}

type PostData struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Data    string            `json:"data"`
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
			log.Println(err)
		}
		headers := make(map[string]string, len(res.Header))
		for k, v := range res.Header {
			headers[k] = v[0]
		}
		var buf bytes.Buffer
		io.Copy(&buf, res.Body)
		resData := PostData{
			Url:     d.Url,
			Method:  d.Method,
			Headers: headers,
			Data:    buf.String(),
		}

		if err := json.NewEncoder(w).Encode(resData); err != nil {
			log.Println(err)
		}
	} else if r.Method == http.MethodGet {
		index, _ := Asset("index.html")
		t, err := template.New("index").Parse(string(index))
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error.", http.StatusInternalServerError)
			return
		}
		tmplData := struct {
			Url string
		}{
			address(),
		}
		if err := t.Execute(w, tmplData); err != nil {
			log.Println(err)
		}
	}
}
