package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"

	"github.com/gin-gonic/gin"
)

func BillList(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|bill-list|"
		p := tables.BillDetail{}
		list, err := p.BillList(ctx.DB, "Pesanan Baru")
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
