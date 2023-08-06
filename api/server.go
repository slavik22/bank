package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/slavik22/bank/db/sqlc"
	"github.com/slavik22/bank/token"
	"github.com/slavik22/bank/util"
)

type Server struct {
	config   util.Config
	store    db.Store
	router   *gin.Engine
	jwtMaker *token.JWTMaker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	jwt, err := token.NewJWTMaker(config.SecretKey)

	if err != nil {
		return nil, fmt.Errorf("Cannot create token maker %w", err)
	}

	server := &Server{
		store:    store,
		jwtMaker: jwt,
		config:   config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	//router.POST("/tokens/renew_access", server.renewAccessToken)

	//authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
