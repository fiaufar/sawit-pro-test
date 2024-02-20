package main

import (
	"net/http"
	"os"

	"github.com/fiaufar/sawit-pro-test/constant"
	"github.com/fiaufar/sawit-pro-test/handler"
	"github.com/fiaufar/sawit-pro-test/infrastructure"
	"github.com/fiaufar/sawit-pro-test/repository"
	"github.com/fiaufar/sawit-pro-test/service"
	"github.com/fiaufar/sawit-pro-test/util"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	server := newServer()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	authGroup := e.Group("auth")
	authGroup.POST("/login", server.Login)
	authGroup.POST("/register", server.Register)

	accountGroup := e.Group("account")
	config := echojwt.Config{
		NewClaimsFunc: util.EchoClaimFunc,
		SigningKey:    []byte(constant.JWT_SECRET_KEY),
	}
	accountGroup.Use(echojwt.WithConfig(config))
	accountGroup.GET("/profile", server.Profile)
	accountGroup.PUT("/profile", server.UpdateProfile)

	// generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	// dbDsn := "postgres://postgres:postgres@localhost:5432/sawit-pro?sslmode=disable"
	dbConnection := infrastructure.NewDbConnection(infrastructure.NewDbConnectionOptions{
		Dsn: dbDsn,
	})
	util.Log.Info(dbConnection)
	// var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
	// 	Dsn: dbDsn,
	// })

	userRepo := repository.NewUserRepository(repository.NewUserRepositoryOptions{
		DbConn: dbConnection,
	})

	userCredentialRepo := repository.NewUserCredentialRepository(repository.NewUserCredentialRepositoryOptions{
		DbConn: dbConnection,
	})
	authService := service.NewAuthService(service.NewAuthServiceOptions{
		UserRepo:           userRepo,
		UserCredentialRepo: userCredentialRepo,
	})
	accountService := service.NewAccountService(service.NewAccountServiceOptions{
		UserRepo: userRepo,
	})

	validator := util.NewValidator()
	responseUtil := util.NewResponseUtil()

	opts := handler.NewServerOptions{
		AuthService:    authService,
		AccountService: accountService,
		Validator:      validator,
		ResponseUtil:   responseUtil,
	}
	return handler.NewServer(opts)
}
