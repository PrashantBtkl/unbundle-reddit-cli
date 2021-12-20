package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func MakeRequest(req *http.Request) ([]byte, error) {
	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Println("Error occured while making request")
		return nil, err
	}

	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)
	return resp_body, nil

}
