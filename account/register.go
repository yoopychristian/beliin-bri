package account

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"strconv"
	"time"

	"beliin-bri/tools"

	"github.com/gin-gonic/gin"
)

func Register(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|register|"
		randNum := tools.RandomNumber(100, 999)
		strRandNum := strconv.Itoa(randNum)
		now := time.Now()
		timeStr := now.Format("0405")
		id_user := timeStr + strRandNum

		input := shared.ParamAccountRegister{}
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
		//Hash Password
		hash, err := h.HashPassword(input.Password)
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "hash-password",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//fullname-rule
		if err := h.FullnameRule(input.Nama); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "fullname-rule",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//username-rule
		if err := h.UsernameRule(input.Username); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "username-rule",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//password-rule
		if err := h.PasswordRule(input.Password); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "password-rule",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//email-rule
		if err := h.EmailRule(input.Email); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "email-rule",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//phone-rule
		if err := h.PhoneRule(input.NoPonsel); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "noponsel-rule",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//Check User
		user := tables.Registration{}
		checkEmail, err := user.EmailExist(ctx.DB, input.Email)
		if err != nil {
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

		if !checkEmail {
			h.BadResponseExist(c, "your email is already registered")
			return
		}

		checkUser, err := user.UsernameExist(ctx.DB, input.Username)
		if err != nil {
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

		if !checkUser {
			h.BadResponseExist(c, "your username is already registered")
			return
		}
		input.IDUser = id_user

		//data logic
		if err := user.Create(ctx.DB, input.IDUser, input.Nama, input.Username, hash, input.Email, input.NoPonsel, input.NoKTP, true); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "result-user",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		//Account Information
		h.GoodResponse(c, shared.ResponseDetail{
			IDUser:   user.IDUser,
			Nama:     user.Nama,
			Username: user.Username,
			Email:    user.Email,
			NoPonsel: user.NoTelepon,
		})
	}
}
