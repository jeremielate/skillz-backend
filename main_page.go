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
		return nil
	}
	headers := make(map[string][]string, len(d.Headers))
	for k, v := range d.Headers {
		headers[k] = []string{v}
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(d.Data)
	req, err := http.NewRequest(method, d.Url, &buf)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = headers
	return req
}

type PostData struct {
	Url     string                 `json:"url"`
	Method  string                 `json:"method"`
	Headers map[string]string      `json:"headers"`
	Data    map[string]interface{} `json:"data"`
	Raw     string                 `json:"raw"`
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var d PostData
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			log.Println(err)
			return
		}
		req := makeRequest(d)
		if req == nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		headers := make(map[string]string, len(res.Header))
		for k, v := range res.Header {
			headers[k] = v[0]
		}
		resData := PostData{
			Url:     d.Url,
			Method:  d.Method,
			Headers: headers,
		}
		var buf bytes.Buffer
		io.Copy(&buf, res.Body)

		resData.Raw = buf.String()
		if err := json.Unmarshal(buf.Bytes(), resData.Data); err != nil {
			resData.Data = make(map[string]interface{}, 0)
			log.Println(err)
		}
		if err := json.NewEncoder(w).Encode(resData); err != nil {
			log.Println(err)
		}
	} else if r.Method == http.MethodGet {
		index, _ := Asset("static/index.html")
		t, err := template.New("index").Parse(string(index))
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error.", http.StatusInternalServerError)
			return
		}
		tmplData := struct {
			Url string
		}{
			"http://" + address(),
		}
		if err := t.Execute(w, tmplData); err != nil {
			log.Println(err)
		}
	}
}
