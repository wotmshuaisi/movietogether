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
	natscon := initNATS()
	httphandlers := handler.RegisterHTTPHandlers(natscon)
	// service part
	initWebService(httphandlers)
}

// init part

func initWebService(router *mux.Router) {
	err := http.ListenAndServeTLS(config.HTTPADDR, config.CRTFILE, config.KEYFILE, router)
	if err != nil {
		logrus.WithError(err).Fatalln("init web service error.")
	}
	fmt.Println("-> http service listening", config.HTTPADDR)
}

func initNATS() *nats.Conn {
	natsaddr := config.NATSADDR
	natsCon, err := nats.Connect(natsaddr)
	if err != nil {
		logrus.Fatalln("nats connect error: %s", err)
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
		logrus.WithError(err).Fatalln("file to write log.")
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return
}
