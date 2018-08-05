package handler

import (
	"github.com/gorilla/websocket"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"text/template"
)

type clients struct {
	wsCon *websocket.Conn
	Name  string
}

var (
	upgrader    = websocket.Upgrader{}
	connections = map[string]*clients{}
)

func readTpl(location string) (*template.Template, error) {
	var tempTpl *template.Template
	fileObj, err := ioutil.ReadFile("html/index.html")
	if err != nil {
		return nil, err
	}
	tempTpl = template.Must(template.New("").Parse(string(fileObj)))
	return tempTpl, nil
}

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var v = struct {
		Title string
	}{
		Title: "test",
	}
	tpl, err := readTpl("html/index.html")
	if err != nil {
		http.Error(w, "No content", 204)
		return
	}
	tpl.Execute(w, &v)
}

// WebsocketHandler ...
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// connection part
	wsCon, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// create connection failed
		if _, ok := err.(websocket.HandshakeError); !ok {
			logrus.Errorln(err)
		}
		return
	}
	// user part
	username := uuid.New()
	defer func() {
		wsCon.Close()
		delete(connections, username)
	}()
	client := clients{
		wsCon: wsCon,
		Name:  username,
	}
	connections[username] = &client
	logrus.Infoln("current online: ", len(connections))
	// message part
	for {
		mt, message, err := wsCon.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, 1001) {
				logrus.Infoln("client lost connection")
				break
			}
			logrus.WithError(err).Infoln("err recv msg")
			break
		}
		if len(message) <= 0 {
			continue
		}
		// sned msg
		tempmsg := []byte(username + ": " + string(message))
		for _, item := range connections {
			item.wsCon.WriteMessage(mt, tempmsg)
		}
	}
}
