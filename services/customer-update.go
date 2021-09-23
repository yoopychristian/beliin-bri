package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateCustomer(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|customer-update|"
		now := time.Now()
		input := shared.ParamCustomerEdit{}
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

		//id-pelanggan
		if err := h.MustNotEmpty(input.IDPelanggan, "id-pelanggan"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "idPelanggan-mustnotempty",
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

		//alamat
		if err := h.MustNotEmpty(input.AlamatPengiriman, "alamat"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "alamat-mustnotempty",
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

		customer := tables.Customer{}
		if err := customer.UpdateCustomer(ctx.DB, input.IDPelanggan, input.Nama, input.Email, input.AlamatPengiriman, input.Kota, input.NoTelepon, now); err != nil {
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
		h.GoodResponse(c, shared.CustomerResponseEdit{
			IDPelanggan:      input.IDPelanggan,
			Nama:             input.Nama,
			Email:            input.Email,
			AlamatPengiriman: input.AlamatPengiriman,
			Kota:             input.Kota,
			NoTelepon:        input.NoTelepon,
		})
	}
}
