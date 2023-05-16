package ws_test

import (
	"github.com/gorilla/websocket"
	"testing"
)

func TestWsReader(t *testing.T) {
	var ws *websocket.Conn

	messageType, p, err := ws.ReadMessage()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v:%v", messageType, string(p))
}
