package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/net/websocket"
)

type Socket struct {
	*revel.Controller
}

func (c Socket) Feed(user string, ws *websocket.Conn) revel.Result {
	return nil
}
