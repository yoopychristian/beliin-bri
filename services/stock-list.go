package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"

	"github.com/gin-gonic/gin"
)

func StockList(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|stock-list|"
		p := tables.Stock{}
		list, err := p.StockList(ctx.DB)
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

		data := []shared.StockListResponse{}
		for _, row := range list {
			data = append(data, shared.StockListResponse{
				NamaBarang:   row.NamaBarang,
				Deskripsi:    row.Deskripsi,
				JumlahBarang: row.JumlahBarang,
				GambarBarang: row.GambarBarang,
			})
		}

		h.GoodResponse(c, data)
	}
}
