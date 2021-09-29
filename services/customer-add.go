package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"strconv"
	"time"

	"beliin-bri/tools"

	"github.com/gin-gonic/gin"
)

func CustomerAdd(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|customer-add|"
		now := time.Now()
		randNum := tools.RandomNumber(100000, 999999)
		strRandNum := strconv.Itoa(randNum)
		input := shared.ParamCustomer{}
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

		//id-user
		if err := h.MustNotEmpty(input.IDUser, "id-user"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "idUser-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//id-va
		if err := h.MustNotEmpty(input.IDVa, "id-va"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "idVA-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//nama
		if err := h.MustNotEmpty(input.Nama, "nama"); err != nil {
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

		//email
		if err := h.MustNotEmpty(input.Email, "email"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "email-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//alamat-oengiriman
		if err := h.MustNotEmpty(input.AlamatPengiriman, "alamat-pengiriman"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "alamat-pengiriman-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//kota
		if err := h.MustNotEmpty(input.Kota, "kota"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "kota-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//no-telepon
		if err := h.MustNotEmpty(input.NoTelepon, "no-telepon"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "noTelepon-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		stock := tables.Customer{}
		if err := stock.Create(ctx.DB, strRandNum, input.IDUser, input.IDVa, input.Nama, input.Email, input.AlamatPengiriman, input.Kota, input.NoTelepon, now, true); err != nil {
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

		//Account Information
		h.GoodResponse(c, shared.CustomerResponse{
			IDPelanggan:      strRandNum,
			IDUser:           input.IDUser,
			IDVa:             input.IDVa,
			Nama:             input.Nama,
			Email:            input.Email,
			AlamatPengiriman: input.AlamatPengiriman,
			Kota:             input.Kota,
			NoTelepon:        input.NoTelepon,
		})
	}
}
