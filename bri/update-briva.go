package bri

import (
	"bytes"
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

func Update(custCode, nama, amount string) {
	token, err := GetToken()
	if err != nil {
		fmt.Println("ERROR:" + err.Error())
	}
	s := "https://sandbox.partner.api.bri.co.id/v1/briva"
	method := "PUT"
	key := []byte("4VPISJ0GmtfqZh8n")

	expiredTime := time.Now().AddDate(0, 0, 2)
	payload := Invoice{
		InstitutionCode: InstitutionCode,
		BrivaNo:         BrivaNo,
		CustCode:        custCode,
		Nama:            nama,
		Amount:          amount,
		Keterangan:      "",
		ExpiredDate:     expiredTime.String(),
	}

	client := &http.Client{}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("ERROR:" + err.Error())
	}

	req, err := http.NewRequest(method, s, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("ERROR:" + err.Error())
	}

	u, err := url.Parse(s)
	if err != nil {
		fmt.Println(err)
	}

	time := (time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))

	path := u.Path
	data := "path=" + path + "&verb=" + method + "&token=" + "Bearer " + token + "&timestamp=" + time + "&body=" + string(payloadJSON)
	fmt.Println(data)
	keyData := []byte(key)
	h := hmac.New(sha256.New, keyData)
	h.Write([]byte(data))
	dataEncrypt := (base64.StdEncoding.EncodeToString(h.Sum(nil)))
	fmt.Println(dataEncrypt)

	req.Header.Add("Content-Type", "application/json")
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

	result := InvoiceResponse{}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.Status, result.Data.CustCode)
}
