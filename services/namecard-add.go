package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"strconv"

	"beliin-bri/tools"

	"github.com/gin-gonic/gin"
)

func NameCardAdd(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|namecard-add|"
		randNum := tools.RandomNumber(100000, 999999)
		strRandNum := strconv.Itoa(randNum)
		input := shared.ParamNameCard{}
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

		//nama-toko
		if err := h.MustNotEmpty(input.NamaToko, "nama-toko"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "nama-toko-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//bidang-usaha
		if err := h.MustNotEmpty(input.BidangUsaha, "bidang-usaha"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "bidang-usaha-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//alamat
		if err := h.MustNotEmpty(input.Alamat, "alamat"); err != nil {
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

		//no-telepon
		if err := h.MustNotEmpty(input.NoTelepon, "no-telepon"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "no-telepon-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		nameCard := tables.NameCard{}
		if err := nameCard.Create(ctx.DB, strRandNum, input.IDUser, input.NamaToko, input.BidangUsaha, input.Alamat, input.NoTelepon, true); err != nil {
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
		h.GoodResponse(c, nil)
	}
}
