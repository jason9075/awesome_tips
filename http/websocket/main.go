package main

import (
	"log"
	"net/http"
    "time"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許所有跨域請求
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        log.Fatal(err)
		return
	}
	defer conn.Close()

	for {
        time.Sleep(5 * time.Second)
        message := "Hello from server! time: " + time.Now().Format("2006-01-02 15:04:05") + "\n"
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
    fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/echo", echo)

    log.Println("http server started on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
