package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/config"
	"github.com/wotmshuaisi/movietogether/handler"
	"net/http"
	"os"
)

func main() {
	// init
	initLogger()
	weblog := initWebLogger()
	natscon := initNATS()
	httphandlers := handler.RegisterHTTPHandlers(natscon, weblog)
	// service part
	initWebService(httphandlers)
}

// init part

func initWebService(router *mux.Router) {
	fmt.Println("-> http service listening", config.HTTPADDR)
	err := http.ListenAndServeTLS(config.HTTPADDR, config.CRTFILE, config.KEYFILE, router)
	if err != nil {
		fmt.Println("init web service error.", err)
		logrus.WithError(err).Fatalln("init web service error.")
		return
	}
}

func initNATS() *nats.Conn {
	natsaddr := config.NATSADDR
	natsCon, err := nats.Connect(natsaddr)
	if err != nil {
		fmt.Println("faile to connect nats.", err)
		logrus.WithError(err).Fatalln("faile to connect nats.")
		return nil
	}
	return natsCon
}

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
