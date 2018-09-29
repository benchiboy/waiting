package util

import (
	"io/ioutil"
	"log"

	"errors"
	"net/http"
	"net/url"
)

var (
	baidu_url string
)

func Http_get(kvMap map[string]string) (string, error) {
	baidu_url = "http://api.map.baidu.com/cloudgc/v1?address=河北元氏县官庄&ak=QjOpH9XPrX9Ak5qhFQYXNE2hP3KMQPms"
	urlVal := url.Values{}
	for k, v := range kvMap {
		urlVal.Add(k, v)
	}
	log.Println("urlVal Encode=", urlVal)
	urlEncode := urlVal.Encode()
	log.Println(baidu_url)
	client := &http.Client{}

	reqest, err := http.NewRequest("GET", baidu_url+urlEncode, nil)
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
