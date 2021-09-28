package main

import (
	"context"
	"crypto/rand"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"beliin-bri/account"
	"beliin-bri/bri"
	cfg "beliin-bri/configuration"
	"beliin-bri/services"

	h "beliin-bri/helpers"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.uber.org/zap"
)

func main() {
	//init context
	ctx, err := h.NewRepositoryContext(rand.Reader, &http.Transport{})
	if err != nil {
		log.Fatal("can't init service context :", err)
	}

	//gin setup
	gin.SetMode(gin.ReleaseMode)

	r := Routing(ctx)
	//http server
	srv := &http.Server{Addr: ":" + ctx.Config.App.Port, Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ctx.Log.Fatal("can't run service", zap.Error(err))
		}
	}()
	ctx.Log.Info(ctx.Config.App.Name + " initiated at port " + ctx.Config.App.Port)

	// gracefully shutdown
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx.Log.Info("Shutdown " + ctx.Config.App.Name + " repository")

	cts, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(cts); err != nil {
		ctx.Log.Warn("can't shutdown "+ctx.Config.App.Name+" repository", zap.Error(err))
	}

	ctx.Log.Info(ctx.Config.App.Name + " repository exiting")
}

func Routing(ctx cfg.RepositoryContext) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	pprof.Register(r)

	p := ginprometheus.NewPrometheus("gin")

	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL
		url.RawQuery = ""
		return url.String()
	}
	p.Use(r)

	//account route
	private := r.Group("/account")
	{
		private.POST("/login", account.Login(ctx))
		private.POST("/register", account.Register(ctx))
	}

	//services
	function := r.Group("/services")
	{
		function.POST("/add-stock", services.AddStock(ctx))
		function.DELETE("/delete-stock", services.DeleteStock(ctx))
		function.GET("/list-stock", services.StockList(ctx))
		function.PUT("/update-stock", services.UpdateStock(ctx))
		function.POST("/add-customer", services.CustomerAdd(ctx))
		function.DELETE("/delete-customer", services.DeleteCustomer(ctx))
		function.GET("/list-customer", services.CustomerList(ctx))
		function.PUT("/update-customer", services.UpdateCustomer(ctx))
		function.POST("/add-order", services.OrderAdd(ctx))
		function.DELETE("/delete-order", services.DeleteOrder(ctx))
		function.GET("/list-order", services.OrderList(ctx))
		function.PUT("/update-order", services.UpdateOrder(ctx))
		function.PUT("/update-order/success", services.OrderSuccess(ctx))
		function.PUT("/update-order/cancel", services.OrderCancel(ctx))
		function.POST("/create-va", bri.Create(ctx))
		function.GET("/bill-detail", services.BillDetail(ctx))
		function.GET("/bill-list", services.BillList(ctx))
		function.POST("/send-bill", services.SendBillDetail(ctx))
		function.GET("/name-card", services.NameCardGet(ctx))
		function.POST("/name-card", services.NameCardAdd(ctx))
		function.PUT("/name-card", services.UpdateNameCard(ctx))
		//function.POST("/get-va", bri.GetBriva(ctx))
	}

	return r
}
