package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var log = logrus.New()

func Ping(c *gin.Context) {
	log.WithFields(logrus.Fields{
		"endpoint": "ping",
		"ts": time.Now().Format("2006-01-02 15:04:05"),
	}).Info("test")

	c.JSON(http.StatusOK, gin.H{
		"name":   "zhangji",
		"locale": "beijing",
		"resp":   "pong",
	})
}

func main() {

	log.SetFormatter(&logrus.JSONFormatter{})
	os.Mkdir("./logs", os.ModePerm)
	file, err := os.OpenFile("./logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		fmt.Println(err)
		log.Info("Failed to log to file, using default stderr")
		return
	}

	router := gin.Default()
	router.GET("/ping", Ping)

	s := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		IdleTimeout:       300 * time.Second,
		WriteTimeout:      20 * time.Second,
	}

	if err := gracehttp.Serve(s); err != nil {
		fmt.Println(err)
	}
}
