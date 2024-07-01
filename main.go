package main

import (
	"exoplant_services/config"
	"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/rs/cors"
	"exoplant_services/pkg/controllers"

	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	configg := config.NewConfig()
	config.PooledConnectDB(configg)
	r := gin.New()
	r.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/exoplanate-ms"
	productGroup := r.Group("/exoplanate-ms")
	controllers.HealthCheckRoutes(productGroup)
	controllers.Exoplanate(productGroup)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tcp_port := configg.PORT
	address := fmt.Sprintf(":%v", tcp_port)
	r.Run(address) // listen and serve on 0.0.0.0:8000
}
