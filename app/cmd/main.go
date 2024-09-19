package main

// Версия 0.1.1 от 06.09.2024

import (
	"github.com/KovalevNikolay710/StockManager/app/internal/server"
	"github.com/KovalevNikolay710/StockManager/app/internal/server/config"
)

func main() {
	cfg := new(config.Config)
	cfg.Read()
	server := server.NewServer(cfg)
	server.Run()
}
