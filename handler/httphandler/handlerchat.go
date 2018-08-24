package httphandler

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"

	"github.com/sirupsen/logrus"
)

func generatetoken() string {
	randomStr := strconv.FormatInt(time.Now().Unix(), 10) + "qGERH$#wg34"
	byetStr := []byte(randomStr)
	byteRandomStr := sha256.Sum256(byetStr)
	result := hex.EncodeToString(byteRandomStr[:])
	return result
}

func jsonerrorreturn(err error, errcode int, w http.ResponseWriter) http.ResponseWriter {
	var res result
	if err != nil {
		res.Err = err.Error()
	} else {
		res.Err = ""
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errcode)
	j, _ := json.Marshal(res)
	w.Write(j)
	return w
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	return
}

// Register regiter user to database
func (handlers *HTTPHandlers) Register(w http.ResponseWriter, r *http.Request) {
	tokenObj, _ := r.Cookie("token")
	if tokenObj != nil {
		_, err := handlers.Model.ClientGet(tokenObj.Value)
		if err == nil {
			w = jsonerrorreturn(nil, 200, w)
			return
		}
	}
	var u user
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &u)
	if err != nil {
		w = jsonerrorreturn(err, 400, w)
		return
	}
	token := generatetoken()
	err = handlers.Model.ClientCreate(token, u.Name)
	if err != nil {
		w = jsonerrorreturn(err, 200, w)
		return
	}
	cookie := http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24 * 360),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(&result{Err: ""})
	w.Write(j)
	return
}

// CheckUser get user by token
func (handlers *HTTPHandlers) CheckUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		w = jsonerrorreturn(err, 400, w)
		return
	}
	c, err := handlers.Model.ClientGet(cookie.Value)
	if err != nil {
		w = jsonerrorreturn(err, 200, w)
		return
	}
	if c == nil || len(c.Token) <= 0 {
		w = jsonerrorreturn(errors.New("client not exists"), 403, w)
		return
	}
	w.Header().Set("status", "200")
	res := &user{
		Name: c.Name,
	}
	j, _ := json.Marshal(res)
	w.Write(j)
	w.Header().Set("Content-Type", "application/json")
	return
}

// Chat interface :  recv message && broadcast
func (handlers *HTTPHandlers) Chat(w http.ResponseWriter, r *http.Request) {
	wsCon, err := handlers.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithError(err).Errorln("http Chat websocket accepting connection")
	}
	name := handlers.Name
	handlers.WsClients[name] = wsCon
	defer func() {
		delete(handlers.WsClients, name)
		wsCon.Close()
		return
	}()
	// websocket connection established
	for {
		_, msg, err := wsCon.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, 1001) || websocket.IsCloseError(err, 1005) {
				logrus.Infoln("http Chat client " + name + " lost connection")
				break
			}
			logrus.WithError(err).Infoln("http Chat client " + name + " lost connection")
			break
		}
		if len(msg) <= 0 {
			continue
		}
		data, err := json.Marshal(&message{
			Datetime: time.Now(),
			Message:  string(msg),
			Name:     name,
		})
		if err != nil {
			logrus.WithError(err).Infoln("http Chat json encode msg error")
			continue
		}
		handlers.MsgQueue <- data
		handlers.Model.MessageCreate(name, string(msg))
	}
	return
}

// History get paged data
func (handlers *HTTPHandlers) History(w http.ResponseWriter, r *http.Request) {
	// wsCon, err :=
	return
}

// Broadcast send msg to client from queue
func (handlers *HTTPHandlers) Broadcast() {
	// websocket connection established
	for {
		jsondata := <-handlers.MsgQueue
		for _, item := range handlers.WsClients {
			err := item.WriteJSON(jsondata)
			if err != nil {
				logrus.WithError(err).Errorln("Broadcast")
			}
		}
	}
}
