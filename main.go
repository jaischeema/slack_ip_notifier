// Borrowed IP detection from https://gist.github.com/jniltinho/9788121
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var slackURL string
	flag.StringVar(&slackURL, "slack_hook_url", "", "incoming webhook integration url")
	flag.Parse()

	if slackURL == "" {
		fmt.Println("Bad value for slack webhook url, use --slack_hook_url to provide proper url")
		os.Exit(1)
	}

	ipAddress, err := GetIPAddress()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	PostToSlack(slackURL, ipAddress)
}

func GetIPAddress() (string, error) {
	var ipAddress []byte
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return string(ipAddress), err
	}
	defer resp.Body.Close()

	ipAddress, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return string(ipAddress), err
	}

	return string(ipAddress), nil
}

func PostToSlack(slackURL string, ipAddress string) {
	payload := map[string]string{"text": "IP Address: " + ipAddress}
	jsonStr, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
}
