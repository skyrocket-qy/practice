package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	fmt.Println("starting server...")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Println("upgrade error: ", err)
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Println("read error: ", err)
				}
				fmt.Printf("Response from server: %s (op=%d)\n", string(msg), op)
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					fmt.Println("write error: ", err)
				}
			}
		}()
	}))
}
