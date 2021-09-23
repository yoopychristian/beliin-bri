package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"

	"github.com/gin-gonic/gin"
)

func OrderList(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|order-list|"
		p := tables.Order{}
		list, err := p.OrderList(ctx.DB)
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

		data := []shared.OrderResponseList{}
		for _, row := range list {
			data = append(data, shared.OrderResponseList{
				IDOrder:           row.IDOrder,
				JumlahBarang:      row.JumlahBarang,
				TotalHarga:        row.TotalHarga,
				PilihanPengiriman: row.PilihanPengiriman,
				OrderStatus:       row.OrderStatus,
			})
		}

		h.GoodResponse(c, data)
	}
}
