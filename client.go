package pokermavensclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func EncodedHttpClientRequest(method string, uri string, parameters map[string]string) (string, error) {
	data := url.Values{}
	for k, v := range parameters {
		data.Set(k, v)
	}
	client := &http.Client{}
	r, err := http.NewRequest(method, uri, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body), err
}
