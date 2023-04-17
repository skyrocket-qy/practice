package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	conn, _, _, err := ws.Dialer{}.Dial(context.TODO(), "ws://localhost:8080")
	if err != nil {
		fmt.Println("upgrade error: ", err)
	}
	defer conn.Close()

	msg := []byte("hello, server!")
	err = wsutil.WriteClientMessage(conn, ws.OpText, msg)
	if err != nil {
		fmt.Println("write error: ", err)
	}

	// read response
	resp, op, err := wsutil.ReadServerData(conn)
	if err != nil {
		fmt.Println("read error: ", err)
	}

	fmt.Printf("Response from server: %s (op=%d)\n", resp, op)
	time.Sleep(10 * time.Second)
}
