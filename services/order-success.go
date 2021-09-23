package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrderSuccess(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|update-order|"
		input := shared.ParamOrderID{}
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

		orders := tables.Order{}
		if err := orders.StatusOrder(ctx.DB, input.IDOrder, "Pesanan Selesai"); err != nil {
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
		c.JSON(http.StatusOK, gin.H{
			"data": "Pesanan Selesai",
		})
	}
}
