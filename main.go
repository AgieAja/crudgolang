package main

import (
	configDb "crudgolang/config"
	deliveryHttp "crudgolang/delivery"
	phoneBookRepo "crudgolang/repository/phonebookrepo"
	phoneBookUC "crudgolang/usecase/phonebookuc"
	"net/http"

	"os"
	"regexp"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	configDb.InitConnMySQLDB()
}

var rxURL = regexp.MustCompile(`^/regexp\d*`)

func main() {
	port := "8001"

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	r.Use(logger.SetLogger())

	// Custom logger
	subLog := zerolog.New(os.Stdout).With().Logger()

	r.Use(logger.SetLogger(logger.Config{
		Logger:         &subLog,
		UTC:            true,
		SkipPath:       []string{"/skip"},
		SkipPathRegexp: rxURL,
	}))

	r.Use(gin.Recovery())

	if port == "" {
		log.Error().Msg("port cant empty")
		return
	}

	db, err := configDb.GetMySQLDB()
	if err != nil {
		log.Error().Msg("Failed Connect to database =" + err.Error())
		return
	}

	defer db.Close()

	phoneRepo := phoneBookRepo.NewPhoneBookRepo(db)
	phoneUC := phoneBookUC.NewPhoneBookUseCase(phoneRepo)

	deliveryHttp.NewPhoneBookHTTPHandler(r, phoneUC)
	log.Info().Msg("Service Running at port : " + port)

	//api for login user
	// r.POST("api/login", authMiddleware.LoginHandler)

	if errHTTP := http.ListenAndServe(":"+port, r); errHTTP != nil {
		log.Error().Msg(errHTTP.Error())
	}
}
