package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SiteCheck home",
	})
	return
}

// эндпойнт функция для получения времени доступа к определенному сайту
func getAccessTime(c *gin.Context) {
	site := c.Query("url")
	if isUrl(site) != true {
		c.String(http.StatusBadRequest, "Неверный формат url! введите url в виде: http://адрессайта.домен")
		fmt.Println("Неверный формат url!")
	} else {
		elapsed := getTime(site)
		c.JSON(http.StatusOK, elapsed)
	}
	// счетчик для статистики
	var r = stat["SiteEndpoint"]
	r = r + 1
	stat["SiteEndpoint"] = r
}

// эндпойнт функция для получения сайта с минимальным временем доступа
func getMinTime(c *gin.Context) {
	var minTimeSite = getMinAccessTimeSite()
	c.JSON(http.StatusOK, minTimeSite)
	// счетчик для статистики
	var r = stat["MinTimeEndpoint"]
	r = r + 1
	stat["MinTimeEndpoint"] = r
}

// эндпойнт функция для получения сайта с максимальным временем доступа
func getMaxTime(c *gin.Context) {
	var maxTimeSite = getMaxAccessTimeSite()
	c.JSON(http.StatusOK, maxTimeSite)
	// счетчик для статистики
	var r = stat["MaxTimeEndpoint"]
	r = r + 1
	stat["MaxTimeEndpoint"] = r
}

// эндпойнт функция для получения сатистики по запросам
func getStat(c *gin.Context) {
	c.JSON(http.StatusOK, stat)
}

// функция проверки url
func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// получение сайта с максимальным временем доступа (за весь период проверок)
func getMaxAccessTimeSite() string {
	var maxTime float64
	for n := range timeSite {
		maxTime = n
		break
	}
	for n := range timeSite {
		if n > maxTime {
			maxTime = n
		}
	}
	var maxTimeSite = timeSite[maxTime]
	return maxTimeSite
}

// получение сайта с минимальным временем доступа (за весь период проверок)
func getMinAccessTimeSite() string {
	var minTime float64
	for n := range timeSite {
		minTime = n
		break
	}
	for n := range timeSite {
		if n < minTime {
			minTime = n
		}
	}
	var minTimeSite = timeSite[minTime]
	return minTimeSite
}
