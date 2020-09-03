package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// 試すよう
func setCookie() {
	url := "http://localhost:8080/show"
	cookie := &http.Cookie{
		Name:  "hoge",
		Value: "hugahuga",
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest", err)
	}
	req.AddCookie(cookie)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("client.Do", err)
	}

	dumpResp, _ := httputil.DumpResponse(res, true)
	fmt.Printf("%s", dumpResp)
}

func main() {
	setCookie()
}
