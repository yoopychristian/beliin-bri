package bri

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)



func GetBriva(custCode string) {
	token, err := GetToken()
	if err != nil {
		fmt.Println("ERROR:" + err.Error())
	}
	s := "https://sandbox.partner.api.bri.co.id/v1/briva/status/" + InstitutionCode + "/" + BrivaNo + "/" + custCode
	method := "GET"

	fmt.Println(s)
	client := &http.Client{}
	req, err := http.NewRequest(method, s, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	u, err := url.Parse(s)
	if err != nil {
		fmt.Println(err)
	}

	time := (time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
	path := u.Path
	data := "path=" + path + "&verb=" + method + "&token=" + "Bearer " + token + "&timestamp=" + time + "&body="
	fmt.Println(data)
	h := hmac.New(sha256.New, []byte("4VPISJ0GmtfqZh8n"))
	h.Write([]byte(data))
	dataEncrypt := (base64.StdEncoding.EncodeToString(h.Sum(nil)))
	fmt.Println(dataEncrypt)

	req.Header.Add("BRI-Timestamp", time)
	req.Header.Add("BRI-Signature", dataEncrypt)
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := GetStatusResponse{}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.Status, result.Data.StatusBayar)
}
