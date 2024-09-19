package server

import (
	"context"
	"fmt"

	config "github.com/KovalevNikolay710/StockManager/app/internal/server/config"
	"github.com/KovalevNikolay710/StockManager/app/internal/storage"
	configdb "github.com/KovalevNikolay710/StockManager/app/internal/storage/config"
	"github.com/KovalevNikolay710/StockManager/app/internal/storage/interfaces"
	"github.com/KovalevNikolay710/StockManager/app/internal/storage/postgresql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

var PgClient interfaces.Storage

func (s *Server) Run() {
	cfgDb := new(configdb.ConfigDb)
	cfgDb.Read()

	pool, err := storage.NewStorageClient(context.TODO(), cfgDb)
	if err != nil {
		fmt.Println(err)
	}
	defer pool.Close()
	PgClient = postgresql.NewDatadase(pool)

	r := gin.Default()
	r.POST("/reservation", s.handleReservation)
	r.PATCH("/reservation", s.handleReservation)
	r.DELETE("/reservation", s.handleReservation)
	r.GET("/store", s.handleGetStoreRemains)

	serverPath := s.cfg.Host + s.cfg.Port
	r.Run(serverPath)
}
