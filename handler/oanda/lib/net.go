package lib

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func Get(_url, token string) ([]byte, int, error) {
	cookieJar, _ := cookiejar.New(nil)
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          50,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       5 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		//Timeout:   time.Duration(5 * time.Second),
		Jar:       cookieJar,
		Transport: transport}
	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		return nil, 0, err
	}
	//defer req.Body.Close()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "close")
	if len(token) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	log.Println(res.Header)

	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return body, res.StatusCode, fmt.Errorf("resonse code: %d", res.StatusCode)
	}

	return body, res.StatusCode, nil

}
