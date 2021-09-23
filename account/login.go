package account

import (
	cfg "beliin-bri/configuration"
	tables "beliin-bri/database"
	h "beliin-bri/helpers"
	shared "beliin-bri/shared"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx cfg.RepositoryContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		process := "|services|login|"
		input := shared.ParamLogin{}
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

		//data logic
		auth := tables.UserRegistration{}
		if err := auth.GetByUsername(ctx.DB, input.Username); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "authentication",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		ap := []byte(auth.Password)
		ip := []byte(input.Password)
		// encrypt with bcrypt encryption
		if err := bcrypt.CompareHashAndPassword(ap, ip); err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.ERROR,
				Section:  process + "compare-hash-password",
				Error:    err,
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = input.Username
		claims["level"] = "application"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			h.BadResponse(h.RespParams{
				Log:      ctx.Log,
				Context:  c,
				Severity: h.DEBUG,
				Section:  process + "token",
				Reason:   err.Error(),
				Input:    input,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": t,
			"data": shared.ResponseDetail{
				IDUser:   auth.IDUser,
				Nama:     auth.Nama,
				Username: auth.Username,
				Email:    auth.Email,
				NoPonsel: auth.NoPonsel,
			},
		})
	}
}
