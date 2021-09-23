package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateOrder(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|update-order|"
		now := time.Now()
		input := shared.ParamOrderEdit{}
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

		//id-order
		if err := h.MustNotEmpty(input.IDOrder, "id-order"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "idOrder-mustnotempty",
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

		//total-harga
		if err := h.MustNotEmpty(input.TotalHarga, "total-harga"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "totalHarga-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//jumlah-barang
		if err := h.MustNotEmpty(input.PilihanPengiriman, "pilihan-pengiriman"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "pilihanPengiriman-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		orders := tables.Order{}
		if err := orders.UpdateOrder(ctx.DB, input.IDOrder, input.TotalHarga, input.PilihanPengiriman, input.JumlahBarang, now); err != nil {
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
		h.GoodResponse(c, shared.OrderResponse{
			IDOrder:           input.IDOrder,
			JumlahBarang:      input.JumlahBarang,
			TotalHarga:        input.TotalHarga,
			PilihanPengiriman: input.PilihanPengiriman,
		})
	}
}
