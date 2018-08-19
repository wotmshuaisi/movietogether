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

func jsonerrorreturn(err error, errcode string, w http.ResponseWriter) {
	var res result
	if err != nil {
		res.Err = err.Error()
		j, _ := json.Marshal(res)
		w.Write(j)
		w.Header().Set("status", errcode)
		return
	}
}

// Register regiter user to database
func (handlers *HTTPHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var u user
	var res result
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &u)
	if err != nil {
		jsonerrorreturn(err, "400", w)
	}
	token := generatetoken()
	err = handlers.Model.ClientCreate(token, u.Name)
	if err != nil {
		jsonerrorreturn(err, "200", w)
	}
	cookie := http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24 * 60),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("status", "200")
	res.Err = ""
	j, _ := json.Marshal(res)
	w.Write(j)
	return
}

// CheckUser get user by token
func (handlers *HTTPHandlers) CheckUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		jsonerrorreturn(err, "400", w)
	}
	c, err := handlers.Model.ClientGet(cookie.Value)
	if err != nil {
		jsonerrorreturn(err, "200", w)
	}
	if c == nil || len(c.Token) <= 0 {
		jsonerrorreturn(errors.New("client not exists"), "403", w)
	}
	w.Header().Set("status", "200")
	res := &user{
		Name: c.Name,
	}
	j, _ := json.Marshal(res)
	w.Write(j)
	return
}

// Chat interface :  recv message && broadcast
func (handlers *HTTPHandlers) Chat(w http.ResponseWriter, r *http.Request) {
	wsCon, err := handlers.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithError(err).Errorln("http Chat websocket accepting connection")
	}
	name := r.Context().Value(namekey).(string)
	defer func() {
		wsCon.Close()
		return
	}()
	// websocket connection established
	for {
		_, msg, err := wsCon.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, 1001) {
				logrus.Infoln("http Chat client " + name + " lost connection")
				break
			}
			logrus.WithError(err).Infoln("http Chat receve message from " + name)
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

// Message send msg to client from queue
func (handlers *HTTPHandlers) Message(w http.ResponseWriter, r *http.Request) {
	wsCon, err := handlers.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithError(err).Errorln("http Message websocket accepting connection")
	}
	defer func() {
		wsCon.Close()
		return
	}()
	name := r.Context().Value(namekey).(string)
	// websocket connection established
	for {
		jsondata := <-handlers.MsgQueue
		err := wsCon.WriteJSON(jsondata)
		if err != nil {
			if websocket.IsCloseError(err, 1001) {
				logrus.Infoln("http Message client " + name + " lost connection")
				break
			}
			logrus.WithError(err).Infoln("http Message write message to " + name)
			break
		}
	}
}

// History get paged data
func (handlers *HTTPHandlers) History(w http.ResponseWriter, r *http.Request) {
	// wsCon, err :=
	return
}
