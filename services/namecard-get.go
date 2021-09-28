package services

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"

	shared "beliin-bri/shared"

	"github.com/gin-gonic/gin"
)

func NameCardGet(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|name-card-get|"
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

		p := tables.NameCard{}
		list, err := p.GetByID(ctx.DB, input.ID)
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

		data := []shared.NameCard{}
		for _, row := range list {
			data = append(data, shared.NameCard{
				IDKartuNama: row.IDKartuNama,
				IDUser:      row.IDUser,
				NamaToko:    row.NamaToko,
				BidangUsaha: row.BidangUsaha,
				Alamat:      row.Alamat,
				NoTelepon:   row.NoTelepon,
			})
		}

		h.GoodResponse(c, data)
	}
}
