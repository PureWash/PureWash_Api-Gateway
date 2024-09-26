package handlers

import (
	"api_gateway/internal/configs"
	"api_gateway/internal/pkg/grpcConn"
	"api_gateway/internal/pkg/logger"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Run()
	Stop()
}

type Handler struct {
	engine   *gin.Engine
	services grpc.IServiceManager
	log      logger.ILogger
	cnf      configs.Config
}

func NewHandler(engine *gin.Engine, services grpc.IServiceManager, log logger.ILogger, cnf configs.Config) *Handler {
	return &Handler{
		engine:   engine,
		services: services,
		log:      log,
		cnf:      cnf,
	}
}

func NewServer(cfg configs.Config) Server {

	loggerLevel := logger.LevelDebug
	switch cfg.Environment {
	case configs.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case configs.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}
	log := logger.NewLogger("hadiya.uz", loggerLevel)
	defer logger.Cleanup(log)

	engine := gin.Default()
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowCredentials = true
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowHeaders = append(defaultConfig.AllowHeaders,
		"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token,"+
			"Authorization, accept, origin, Cache-Control, X-Requested-With")
	defaultConfig.AllowHeaders = append(defaultConfig.AllowHeaders, "*")
	defaultConfig.AllowMethods = append(defaultConfig.AllowMethods, "OPTIONS")
	engine.Use(cors.New(defaultConfig))

	service, err := grpc.NewGrpcClients(cfg, log)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	handler := NewHandler(engine, service, log, cfg)
	setUpApi(handler)

	return handler
}

// @title PURE-WASH.UZ App
// @description This API contains the source for the pure_wash.uz app
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /v1

// Run initializes http server
func (h *Handler) Run() {
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL(fmt.Sprintf(
			"%s/%d/swagger/docs.json",
			h.cnf.HTTPHost,
			h.cnf.HTTPPort,
		)),
		ginSwagger.DefaultModelsExpandDepth(-1),
	)
	h.log.Info("server is running: ", logger.Any("HOST", h.cnf.HTTPHost), logger.Any("PORT", h.cnf.HTTPPort))

	if err := h.engine.Run(fmt.Sprintf("%s:%s", h.cnf.HTTPHost, h.cnf.HTTPPort)); err != nil {
		h.log.Error("failed to run server", logger.Error(err))
	}

}

func (h *Handler) Stop() {
	h.log.Info("shutting down")
}
