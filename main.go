package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	client := &http.Client{}
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}
	fmt.Println(url)
	req, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	w.Header().Set("content-type", resp.Header.Get("content-type"))
	body, err := ioutil.ReadAll(resp.Body)
	w.Write(body)
}

func setupCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
