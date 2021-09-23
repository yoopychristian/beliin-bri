package bri

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	"beliin-bri/shared"
	"beliin-bri/tools"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var InstitutionCode = "J104408"
var BrivaNo = "88888"

func Create(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|create-add|"
		now := time.Now()
		randNum := tools.RandomNumber(100000, 999999)
		strRandNum := strconv.Itoa(randNum)
		input := shared.ParamVirtualAccount{}
		if err := c.Bind(&input); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "bind",
				Reason:   "missing input",
			})
			return
		}

		if err := h.MustNotEmpty(input.IDPelanggan, "id-pelanggan"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "idpelanggan-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		if err := h.MustNotEmpty(input.NamaRekening, "nama"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "nama-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		if err := h.MustNotEmpty(input.Amount, "amount"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "amount-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		if err := h.MustNotEmpty(input.NoVa, "no_va"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "no_va-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		token, err := GetToken()
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "get-token",
				Error:    err,
				Reason:   err.Error(),
				Input:    token,
			})
			return
		}

		s := "https://sandbox.partner.api.bri.co.id/v1/briva"
		method := "POST"
		key := []byte("4VPISJ0GmtfqZh8n")

		expiredTime := time.Now().AddDate(0, 0, 2).Format("2006-01-02 15:04:05")
		payload := Invoice{
			InstitutionCode: InstitutionCode,
			BrivaNo:         BrivaNo,
			CustCode:        input.NoVa,
			Nama:            input.NamaRekening,
			Amount:          input.Amount,
			Keterangan:      "",
			ExpiredDate:     expiredTime,
		}

		client := &http.Client{}
		payloadJSON, err := json.Marshal(payload)
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "payload marshal",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		req, err := http.NewRequest(method, s, bytes.NewBuffer(payloadJSON))
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "req-payload",
				Error:    err,
				Reason:   err.Error(),
				Input:    req,
			})
			return
		}

		u, err := url.Parse(s)
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "parse-url",
				Reason:   err.Error(),
				Input:    u,
			})
			return
		}

		time := (time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))

		path := u.Path
		data := "path=" + path + "&verb=" + method + "&token=" + "Bearer " + token + "&timestamp=" + time + "&body=" + string(payloadJSON)
		keyData := []byte(key)
		hmac := hmac.New(sha256.New, keyData)
		hmac.Write([]byte(data))
		dataEncrypt := (base64.StdEncoding.EncodeToString(hmac.Sum(nil)))

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("BRI-Timestamp", time)
		req.Header.Add("BRI-Signature", dataEncrypt)
		req.Header.Add("Authorization", "Bearer "+token)

		res, err := client.Do(req)
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "result",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "read-body",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		fmt.Println(string(body))
		result := InvoiceResponse{}
		if err := json.Unmarshal([]byte(body), &result); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "unmarshal-result",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		va := tables.VirtualAccount{}
		if err := va.Create(ctx.DB, strRandNum, input.IDPelanggan, result.Data.Nama, result.Data.CustCode, "Rupiah", "N", now, true); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "table-add-result",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}
		//VA Information
		h.GoodResponse(c, shared.VAResponse{
			IDUser:       input.IDPelanggan,
			NoVA:         result.Data.CustCode,
			NamaRekening: result.Data.Nama,
			Amount:       result.Data.Amount,
			Status:       "N",
			ExpiredDate:  result.Data.ExpiredDate,
		})

	}

}
