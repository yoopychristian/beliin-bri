package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"

	"github.com/gin-gonic/gin"
)

func BillDetail(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|bill-detail|"
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

		list, err := p.DetailList(ctx.DB, input.ID)
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
				PilihanPengiriman: row.PilihanPengiriman,
				NoVa:              row.NoVa,
			})
		}

		h.GoodResponse(c, data)
	}
}
