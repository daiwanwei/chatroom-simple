package main

import (
	"backend-go/chatroom"
	"fmt"
	"net/http"
)

func main() {
	room := chatroom.NewChatRoom()
	http.Handle("/room", room)
	go room.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	port := ":8080"
	fmt.Printf("server serving: http://localhost%s/\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

}
