package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteStock(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|delete-stock|"
		input := shared.ParamDeleteStock{}
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

		stock := tables.Stock{}
		if err := stock.DeleteStock(ctx.DB, input.IDStock); err != nil {
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
			"data": "delete stock succesfull",
		})
	}
}
