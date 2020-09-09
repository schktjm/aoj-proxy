package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
}

func showCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	for _, c := range r.Cookies() {
		fmt.Println(c)
	}
}

func makeRequest(r *http.Request, url string) (*http.Response, error) {
	client := &http.Client{}

	defer r.Body.Close()
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	for k, vv := range r.Header {
		for _, v := range vv {
			req.Header.Set(k, v)
		}
	}

	cookies := r.Cookies()
	if len(cookies) > 0 {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}

	return client.Do(req)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var url string
	if path := strings.Split(r.URL.Path, "/"); path[1] == "dat" {
		url = fmt.Sprintf("%s/%s", os.Getenv("DAT_BASE_URL"), strings.Join(path[2:], "/"))
	} else {
		url = fmt.Sprintf("%s/%s", os.Getenv("AOJ_BASE_URL"), strings.Join(path[2:], "/"))
	}

	defer r.Body.Close()
	res, err := makeRequest(r, url)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range res.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Fprint(w, string(body))
	fmt.Println("url", url)
}

func main() {
	envLoad()

	fmt.Println("listening localhost:8080...")
	http.HandleFunc("/show", showCookie)
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
