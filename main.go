package main

import (
	"exoplant_services/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	configg := config.NewConfig()
	config.PooledConnectDB(configg)
	r := gin.New()

	tcp_port := configg.PORT
	address := fmt.Sprintf(":%v", tcp_port)
	r.Run(address) // listen and serve on 0.0.0.0:8000
}
