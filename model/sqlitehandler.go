package model

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (handler *dbhandler) ClientGet(token string) (*Clients, error) {
	result := &Clients{}
	sqlstr := "SELECT client_name, client_token FROM 'clients' WHERE client_token = ?"
	err := handler.Db.Get(result, sqlstr, token)
	if err != nil {
		logrus.WithError(err).Errorln("ClientGet")
	}
	return result, err
}

func (handler *dbhandler) ClientCreate(token string, name string) error {
	sqlstr := "INSERT INTO 'clients' ( 'client_name', 'client_token')  VALUES ( ?, ? )"
	_, err := handler.Db.Exec(sqlstr, name, token)
	if err != nil {
		logrus.WithError(err).Errorln("ClientCreate")
	}
	return err
}

func (handler *dbhandler) MessagesSelect(pagenum int, pagesize int) (*MessagesQueryset, error) {
	var count int
	var result []*Messages
	sqlstr := "SELECT COUNT(*) FROM 'messages'"
	err := handler.Db.Get(&count, sqlstr)
	if err != nil {
		logrus.WithError(err).Errorln("MessagesSelect 1")
		return nil, err
	}
	sqlstr = "SELECT datetime, message, name FROM messages LIMIT ?,?"
	err = handler.Db.Select(&result, sqlstr, (pagenum-1)*pagesize, pagesize)
	if err != nil {
		logrus.WithError(err).Errorln("MessagesSelect 2")
		return nil, err
	}
	queryset := &MessagesQueryset{
		TotalNums: count,
		TotalPage: (count + pagesize - 1) / pagesize,
		Page:      pagenum,
		Result:    result,
	}
	return queryset, nil
}

func (handler *dbhandler) MessageCreate(name string, message string) error {
	sqlstr := "INSERT INTO 'messages' ( 'datetime', 'message', 'name')  VALUES ( ?, ?, ? )"
	_, err := handler.Db.Exec(sqlstr, time.Now(), message, name)
	if err != nil {
		logrus.WithError(err).Errorln("MessageCreate")
	}
	return err
}
