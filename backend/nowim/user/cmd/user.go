package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"nowim.user/internal/api"
	_ "nowim.user/internal/config"
	_ "nowim.user/internal/db"
)

func main() {
	e := gin.New()
	api.SetUpHandlers(e)
	if err := e.Run(); err != nil {
		log.Fatalf("server start failed, err: %+v", err)
	}
}
