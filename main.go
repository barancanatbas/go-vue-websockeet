package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/socket", WsEdnPoint)

	log.Fatal(http.ListenAndServe(":9100", router))
}

type Message struct {
	Greeting string `json:"greeting"`
	Sender   int    `json:"sender"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var keys = []string{"selam baran", "iyidir senden naber", "sen kimsin ?", "aynen", "yok daha neler", "güzel bir örnek oldu", "go güzel dil", "fakat random sayı yok", "güzel bir hafta sonundan selamlar", "khanjer"}

func WsEdnPoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("hata: ", err.Error())
		return
	}
	defer wsConn.Close()

	// event loop
	for {
		var msg Message

		// gelen msg değerini json olarak okur ve verilen struct değerine bind eder.
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("error reading JSON : ", err.Error())
			break
		}

		fmt.Println("Alınan mesaj : ", msg.Greeting)

		randnumber := rand.Intn(10)

		newmsg := Message{
			Sender:   2,
			Greeting: keys[randnumber],
		}

		SendMsg(newmsg, wsConn)
	}
}

func SendMsg(msg Message, wsConn *websocket.Conn) {
	err := wsConn.WriteJSON(&msg)
	if err != nil {
		fmt.Println("send msg error : ", err.Error())
	}
}
