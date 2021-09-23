package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"

	"github.com/gin-gonic/gin"
)

func CustomerList(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|customer-list|"
		p := tables.Customer{}
		list, err := p.CustomerList(ctx.DB)
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

		data := []shared.CustomerListResponse{}
		for _, row := range list {
			data = append(data, shared.CustomerListResponse{
				IDPelanggan:      row.IDPelanggan,
				Nama:             row.Nama,
				Email:            row.Email,
				AlamatPengiriman: row.AlamatPengiriman,
				Kota:             row.Kota,
				NoTelepon:        row.NoTelepon,
			})
		}

		h.GoodResponse(c, data)
	}
}
