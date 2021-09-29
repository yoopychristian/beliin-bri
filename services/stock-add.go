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

func AddStock(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|add-stock|"
		now := time.Now()
		randNum := tools.RandomNumber(100000, 999999)
		strRandNum := strconv.Itoa(randNum)
		input := shared.ParamStock{}
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

		//nama-barang
		if err := h.MustNotEmpty(input.NamaBarang, "nama-barang"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "namaBarang-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//deskripsi
		if err := h.MustNotEmpty(input.Deskripsi, "deskripsi"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "deksripsi-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//harga-satuan
		if err := h.MustNotEmpty(input.HargaSatuan, "harga-satuan"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "harga-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//jumlah-barang
		if err := h.NotZero(input.JumlahBarang, "jumlah-barang"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "jumlahBarang-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//gambar-barang
		if err := h.MustNotEmpty(input.GambarBarang, "gambar-barang"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "gambarBarang-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		stock := tables.Stock{}
		if err := stock.Create(ctx.DB, strRandNum, input.IDUser, input.NamaBarang, input.Deskripsi, input.HargaSatuan, input.GambarBarang, input.JumlahBarang, now, true); err != nil {
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
		h.GoodResponse(c, shared.StockReponse{
			IDStock:      strRandNum,
			IDUser:       input.IDUser,
			NamaBarang:   input.NamaBarang,
			Deskripsi:    input.Deskripsi,
			HargaSatuan:  input.HargaSatuan,
			JumlahBarang: input.JumlahBarang,
			GambarBarang: input.GambarBarang,
		})
	}
}
