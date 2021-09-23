package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendBillDetail(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|send-bill-detail|"
		p := tables.BillDetail{}
		input := shared.ParamID{}
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

		if err := h.MustNotEmpty(input.ID, "id"); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "id-mustnotempty",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		list, err := p.SendBill(ctx.DB, input.ID)
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "bind",
				Reason:   "missing input",
			})
			return
		}

		data := []shared.BillDetail{}
		for _, row := range *list {
			data = append(data, shared.BillDetail{
				IDPelanggan:       row.IDPelanggan,
				Nama:              row.Nama,
				Alamat:            row.Alamat,
				NoTelepon:         row.NoTelepon,
				JumlahBarang:      row.JumlahBarang,
				TotalHarga:        row.TotalHarga,
				Email:             row.Email,
				PilihanPengiriman: row.PilihanPengiriman,
				NoVa:              row.NoVa,
			})

		}

		e, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		go h.SendMail("yoopychs@gmail.com", string(e))
		//Account Information
		c.JSON(http.StatusOK, gin.H{
			"data": "Send Invoice Success",
		})

		h.GoodResponse(c, data)
	}
}
