package util

import (
	"io/ioutil"
	"log"

	"errors"
	"net/http"
	"net/url"
	"time"
)

var (
	baidu_url string
)

func Http_get(goUrl string, kvMap map[string]string) (string, error) {
	urlVal := ""
	for k, v := range kvMap {
		urlVal += k
		urlVal += "=" + v
	}
	t := time.Now()
	client := &http.Client{}

	reqest, err := http.NewRequest("GET", goUrl+urlVal, nil)
	if err != nil {
		log.Println("NewRequest is error")
		return "", nil
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		log.Println("client.Do", err)
		return "", nil
	}
	defer response.Body.Close()
	var body []byte

	log.Println("Elapsed:", time.Since(t))
	if response.StatusCode == 200 {
		switch response.Header.Get("Content-Encoding") {
		default:
			bodyByte, _ := ioutil.ReadAll(response.Body)
			body = bodyByte
			log.Println(string(body))
		}
	} else {
		log.Println("http resp code is not 200", err)
		return "", errors.New("http resp code is not 200")
	}
	return "", nil
}

func Http_post(goUrl string, kvMap map[string]string) (string, error) {
	urlVal := ""
	for k, v := range kvMap {
		urlVal += k
		urlVal += "=" + v
	}
	t := time.Now()
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", goUrl+urlVal, nil)
	if err != nil {
		log.Println("NewRequest is error")
		return "", nil
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		log.Println("client.Do", err)
		return "", nil
	}
	defer response.Body.Close()
	var body []byte

	log.Println("Elapsed:", time.Since(t))
	if response.StatusCode == 200 {
		switch response.Header.Get("Content-Encoding") {
		default:
			bodyByte, _ := ioutil.ReadAll(response.Body)
			body = bodyByte
			log.Println(string(body))
		}
	} else {
		log.Println("http resp code is not 200", err)
		return "", errors.New("http resp code is not 200")
	}
	return "", nil
}

func Http_get2(goUrl string, kvMap map[string]string) (string, error) {
	urlVal := url.Values{}
	for k, v := range kvMap {
		urlVal.Add(k, v)

	}
	params := urlVal.Encode()
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", goUrl+params, nil)
	if err != nil {
		log.Println("NewRequest is error")
		return "", nil
	}
	//	log.Println("get url=====>", goUrl+params)
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		log.Println("client.Do", err)
		return "", nil
	}
	defer response.Body.Close()
	var body []byte
	if response.StatusCode == 200 {
		switch response.Header.Get("Content-Encoding") {
		default:
			bodyByte, _ := ioutil.ReadAll(response.Body)
			body = bodyByte
			log.Println(string(body))
		}
	} else {
		log.Println("http resp code is not 200", err)
		return "", errors.New("http resp code is not 200")
	}
	return "", nil
}
