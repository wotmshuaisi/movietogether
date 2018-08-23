package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/websocket"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/config"
	"github.com/wotmshuaisi/movietogether/handler"
	"github.com/wotmshuaisi/movietogether/model"
)

func main() {
	// init
	initLogger()
	// web logrus
	weblog := initWebLogger()
	// rtmp channel
	channel := pubsub.NewQueue()
	// message channel
	msgchannel := make(chan []byte, 2048)
	// websocket
	wsupgrader := &websocket.Upgrader{}
	wsupgrader.HandshakeTimeout = time.Second * 30
	// database
	db := initDB()
	// db model
	m := model.NewModel(db)

	httphandlers := handler.RegisterHTTPHandlers(weblog, channel, wsupgrader, msgchannel, m)
	rtmphandlers := handler.RegisterRTMPHandlers(channel)
	// service part
	go webService(httphandlers)
	rtmpService(rtmphandlers)
}

// services part

func webService(router *mux.Router) {
	fmt.Println("-> http service listening", config.HTTPADDR)
	err := http.ListenAndServe(config.HTTPADDR, router)
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

func initDB() *sqlx.DB {
	var err error
	dbPath := "../data/db.sqlite3?charset=utf8mb4&collation=utf8mb4_unicode_ci"
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("open db [" + dbPath + "] failed")
		logrus.WithError(err).Fatalln("err database connection")
	}
	return db
}

func initWebLogger() *logrus.Logger {
	log := logrus.New()
	var logFilePath = "log/"
	logFile, err := os.OpenFile(logFilePath+"web.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = logFile
	} else {
		fmt.Println("faile to write log.", err)
		logrus.WithError(err).Fatalln("faile to write log.")
	}
	log.Formatter = &logrus.JSONFormatter{}
	return log
}
