package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCustomer(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|customer-delete|"
		input := shared.ParamDeleteCustomer{}
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

		customer := tables.Customer{}
		if err := customer.DeleteCustomer(ctx.DB, input.IDPelanggan); err != nil {
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
			"data": "delete customer succesfull",
		})
	}
}
