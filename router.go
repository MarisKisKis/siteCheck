package main

import (
	"github.com/gin-gonic/gin"
)

// эндпойнты для запросов
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", home)
	//эндпоинт для получения времени доступа к определенному сайту, название сайта передается в параметре ("key") url,
	//пример: http://localhost:8080/time/site?url=http//alipay.com
	r.GET("/time/site", getAccessTime)
	// получение сайта с минимальным временем доступа
	r.GET("/time/min", getMinTime)
	// получение сайта с максимальным временем доступа
	r.GET("/time/max", getMaxTime)
	// получение статистики по запросам
	r.GET("/admin/stat", getStat)
	return r
}
