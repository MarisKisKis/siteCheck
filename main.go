package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var timer int = 1
var checkHistory []string
var histLength int
var timeSite = make(map[float64]string)
var stat = make(map[string]int)

func main() {
	stat["SiteEndpoint"] = 0
	stat["MinTimeEndpoint"] = 0
	stat["MaxTimeEndpoint"] = 0
	sites := getSitesFromFile("sites.txt")
	ch := make(chan byte, 1)
	// цикл, запускающий поочередную проверку сайтов 1 раз в минуту
	for _, site := range sites {
		go func(site string) {
			for {
				tm := time.Now().Format("2006-01-02 15:04:05")
				fmt.Println("Проверяем адрес ", site)
				var checkSite = "http://" + site
				_, msg := check(checkSite)
				t := getTime(checkSite)
				timeSite[t] = site
				logToFile(tm, msg)
				saveHistory(tm, msg)
				time.Sleep(time.Duration(timer) * time.Minute)
			}
			ch <- 1
		}(site)
		fmt.Println(timeSite)
	}
	//запуск сервера
	r := SetupRouter()
	r.Run("localhost:8080")
	<-ch
}

// получение списка сайтов из файла
func getSitesFromFile(namefile string) []string {
	data, err := ioutil.ReadFile(namefile)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(2)
	}
	result := string(data)

	test := strings.Split(result, "\n")
	return test
}

// check проверяет доступность и время для подключения к сайту
func check(url string) (bool, string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		return false, fmt.Sprintf("Ошибка соединения. %s", err)
	}
	defer resp.Body.Close()
	elapsed := time.Since(start).Seconds()
	return true, fmt.Sprintf("Онлайн. http-статус: %d. Время доступа к сайту %s: %v секунд", resp.StatusCode, url, elapsed)
}

// getTime получает время доступа к сайту
func getTime(site string) float64 {
	start := time.Now()
	resp, err := http.Get(site)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	elapsed := time.Since(start).Seconds()
	return elapsed
}

// logToFile сохраняет сообщения в файл
func logToFile(tm, s string) {
	f, err := os.OpenFile("site_check.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(tm, err)
		return
	}
	if _, err = f.WriteString(fmt.Sprintln(tm, s)); err != nil {
		fmt.Println(tm, err)
	}
}

// saveHistory добавляет запись в массив с историей проверок
func saveHistory(tm, s string) {
	checkHistory = append(checkHistory, fmt.Sprintf("%s %s", tm, s))
	if len(checkHistory) > histLength {
		checkHistory = checkHistory[1:]
	}
}
