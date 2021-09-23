package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateStock(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|update-stock|"
		now := time.Now()
		input := shared.ParamStockEdit{}
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

		//id-stock
		if err := h.MustNotEmpty(input.IDStock, "id-stock"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "idStock-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//nama barang
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

		//deksripsi
		if err := h.MustNotEmpty(input.Deskripsi, "deskripsi"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "deskripsi-mustnotempty",
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

		//harga-satuan
		if err := h.MustNotEmpty(input.HargaSatuan, "harga-satuan"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "hargaSatuan-mustnotempty",
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
		if err := stock.UpdateStock(ctx.DB, input.IDStock, input.NamaBarang, input.Deskripsi, input.JumlahBarang, input.HargaSatuan, input.GambarBarang, now); err != nil {
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
			IDStock:      input.IDStock,
			NamaBarang:   input.NamaBarang,
			Deskripsi:    input.Deskripsi,
			JumlahBarang: input.JumlahBarang,
			GambarBarang: input.GambarBarang,
		})
	}
}
