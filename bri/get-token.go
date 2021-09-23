package bri

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var token string
var tokenExpired time.Time

func GetToken() (string, error) {
	errHandle := func(err error) (string, error) {
		return "", err
	}

	result := GetTokenResponse{}

	url := "https://sandbox.partner.api.bri.co.id/oauth/client_credential/accesstoken?grant_type=client_credentials"
	method := "POST"

	payload := strings.NewReader("client_id=DjWdC3ZFIAbhfj2mKaKz4NFKpGeQkGPn&client_secret=4VPISJ0GmtfqZh8n")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return errHandle(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return errHandle(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errHandle(err)
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return errHandle(err)
	}
	token = result.AccessToken

	return token, nil
}

// func needNewToken() bool {
// 	if token == "" {
// 		return true
// 	}

// 	now := time.Now()
// 	diff := tokenExpired.Sub(now)
// 	// If almost expired, need new token
// 	if diff.Minutes() < 5 {
// 		return true
// 	}

// 	return false
// }
