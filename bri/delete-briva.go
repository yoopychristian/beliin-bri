package bri

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Delete() {

	url := "https://sandbox.partner.api.bri.co.id/v1/briva"
	method := "DELETE"

	payload := strings.NewReader(`institutionCode=J104408&brivaNo=77777&custCode=20210818`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("BRI-Timestamp", "2021-09-15T05:34:06.307Z")
	req.Header.Add("BRI-Signature", "bwl8ZIP5sCdcdfNc00tEr8F5m1w/NmgmKV7q9owb68k=")
	req.Header.Add("Authorization", "Bearer ExLqkwRxhfFcdZEyLvggB49cjiWA")
	req.Header.Add("Content-Type", "text/plain")

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
	fmt.Println(string(body))
}
