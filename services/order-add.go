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

func OrderAdd(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|order-add|"
		now := time.Now()
		randNum := tools.RandomNumber(100000, 999999)
		strRandNum := strconv.Itoa(randNum)
		input := shared.ParamOrder{}
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
		if err := h.MustNotEmpty(input.IDStock, "id-va"); err != nil {
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

		//jumlah-barang
		if err := h.NotZero(input.JumlahBarang, "jumlah-barang"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "jumlahBarang-mustnotzero",
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

		//pilihan-pengiriman
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

		input.OrderStatus = "Pesanan Baru"

		orders := tables.Order{}
		if err := orders.Create(ctx.DB, strRandNum, input.IDUser, input.IDStock, input.IDPelanggan, input.TotalHarga, input.PilihanPengiriman, input.OrderStatus, now, input.JumlahBarang, true); err != nil {
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
			IDUser:            input.IDUser,
			IDStock:           input.IDStock,
			IDPelanggan:       input.IDPelanggan,
			JumlahBarang:      input.JumlahBarang,
			TotalHarga:        input.TotalHarga,
			PilihanPengiriman: input.PilihanPengiriman,
			OrderStatus:       input.OrderStatus,
		})
	}
}
