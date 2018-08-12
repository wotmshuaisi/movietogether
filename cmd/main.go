package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/config"
	"github.com/wotmshuaisi/movietogether/handler"
)

func main() {
	// init
	channel := pubsub.NewQueue()
	initLogger()
	weblog := initWebLogger()
	httphandlers := handler.RegisterHTTPHandlers(weblog, channel)
	rtmphandlers := handler.RegisterRTMPHandlers(channel)
	// service part
	go webService(httphandlers)
	rtmpService(rtmphandlers)
}

// services part

func webService(router *mux.Router) {
	fmt.Println("-> http service listening", config.HTTPADDR)
	err := http.ListenAndServeTLS(config.HTTPADDR, config.CRTFILE, config.KEYFILE, router)
	if err != nil {
		fmt.Println("init web service error.", err)
		logrus.WithError(err).Fatalln("init web service error.")
		return
	}
}

func rtmpService(rtmpserver *rtmp.Server) {
	fmt.Println("-> rtmp service listening")
	err := rtmpserver.ListenAndServe()
	if err != nil {
		fmt.Println("init rtmp service error.", err)
		logrus.WithError(err).Fatalln("init rtmp service error.")
		return
	}
}

// init part

func initLogger() {
	var logFilePath = "log/"

	logFile, err := os.OpenFile(logFilePath+"main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		fmt.Println("faile to write log.", err)
		logrus.WithError(err).Fatalln("faile to write log.")
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return
}

func initWebLogger() *logrus.Logger {
	log := logrus.New()
	var logFilePath = "log/"
	logFile, err := os.OpenFile(logFilePath+"web.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = logFile
		// logrus.SetOutput(logFile)
	} else {
		fmt.Println("faile to write log.", err)
		log.WithError(err).Fatalln("faile to write log.")
	}
	log.Formatter = &logrus.JSONFormatter{}
	return log
}
