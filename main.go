package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"
	"github.com/thiagoluiznunes/ze-challenge/data"
	"github.com/thiagoluiznunes/ze-challenge/domain/service"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"github.com/thiagoluiznunes/ze-challenge/server"
	"github.com/thiagoluiznunes/ze-challenge/server/router"
	"github.com/thiagoluiznunes/ze-challenge/server/router/partnerroute"
	"github.com/thiagoluiznunes/ze-challenge/server/serverconfig"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		endAsErr(err, "couldn't read config file")
	}
	_, err = json.Marshal(cfg)
	endAsErr(err, "couldn't marshal config file")

	log.Info(fmt.Sprintf("connecting to the database at %s:%d.", cfg.DBHost, cfg.DBPort))
	db, err := data.Connect(*(cfg))
	endAsErr(err, "couldn't connect to database.")

	err = db.SetIndexes()
	endAsErr(err, "couldn't create to database indexes")

	atInterruption(func() {
		log.Printf("closing database connection.")
		db.Close()
	})

	log.Info(fmt.Sprintf("connecting to new relic at application %s. (%v)", cfg.NewRelicApplicationName, cfg.NewRelicEnabled))
	nrl, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.NewRelicApplicationName),
		newrelic.ConfigLicense(cfg.NewRelicLicenseKey),
		func(nrlConfig *newrelic.Config) {
			nrlConfig.Enabled = cfg.NewRelicEnabled
			nrlConfig.ErrorCollector.Enabled = cfg.NewRelicErrorCollecting
			nrlConfig.DistributedTracer.Enabled = cfg.DistributedTracerEnabled
		},
	)
	endAsErr(err, "couldn't start new relic application.")
	if err := nrl.WaitForConnection(15 * time.Second); nil != err {
		log.Info(fmt.Sprintf("couldn't connect to new relic application: %v", err))
	}

	svc, err := service.New(db, cfg)
	endAsErr(err, "couldn't create service structure.")

	partnerService := service.NewPartnerService(svc)

	initServer(cfg, nrl, partnerService)
}

func initServer(cfg *config.Config, nrl *newrelic.Application, partnerService *service.PartnerService) {

	e := echo.New()

	// Add controllers
	partnerController := partnerroute.NewController(cfg, partnerService)

	// Initialize Routers
	partnerRoute := partnerroute.NewRouter("partner", partnerController)

	srv := server.Instance(e, cfg)

	appRouter := router.New(srv.Echo, cfg, "ze-delivery")
	appRouter.AddRouters(partnerRoute)

	srv.AddAppRouter(appRouter)
	srv.AddMiddleware(middleware.Logger())
	srv.AddMiddleware(nrecho.Middleware(nrl))
	srv.AddMiddleware(serverconfig.NewRelicMiddleware(nrl))

	log.Info("runninng server at localhost:", cfg.HTTPPort)

	err := srv.Run()
	endAsErr(err, "couldn't start the server.")
}

func atInterruption(fn func()) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fn()
		os.Exit(0)
	}()
}

func endAsErr(err error, message string) {
	if err != nil {
		log.Error(message)
		log.Error(err)
		time.Sleep(time.Millisecond * 50) // needed for printing all messages before exiting
		os.Exit(1)
	}
}
