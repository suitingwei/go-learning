package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"net"
	"net/http"
)

type Client struct {
	IsOnline   bool
	Ip         string
	Connection net.Conn
}

var connections map[net.Conn]*Client

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	connections = make(map[net.Conn]*Client)

	go statServerStart()

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Printf("Failed to accept connection:%s", err.Error())
			continue
		}
		client := &Client{
			Ip:         conn.RemoteAddr().String(),
			IsOnline:   true,
			Connection: conn,
		}
		connections[conn] = client

		welcomeMsg := fmt.Sprintf("[%s] has joined the chat room\n", client.Ip)
		broadcastAll(welcomeMsg, nil)

		go handleConnection(client)
	}
}

func statServerStart() {
	http.HandleFunc("/stat", func(writer http.ResponseWriter, request *http.Request) {
		stat := fmt.Sprintf("Connnections:%d", len(connections))
		_, _ = writer.Write([]byte(stat))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func broadcastAll(msg string, exceptClient *Client) {

	for curCon, curClient := range connections {
		if exceptClient == nil || (curCon != exceptClient.Connection) {
			_, err := curCon.Write([]byte(msg))

			if err != nil {
				log.Fatalf("Failed to send data to connection: %s,reason: %s", curClient.Ip, err.Error())
			}
		}
	}
}

func handleConnection(client *Client) {

	for {
		reader := bufio.NewReader(client.Connection)

		msg, err := reader.ReadString('\n')

		if err != nil {
			//检测到EOF，删除这个链接,广播退出消息
			if err == io.EOF {
				delete(connections, client.Connection)
				leaveMsg := fmt.Sprintf("[%s] has leaved the group", client.Ip)
				broadcastAll(leaveMsg, client)
				return
			} else {
				log.Fatalf("Failed to read msg from connection:%s", err.Error())
			}
		}

		log.Printf("Read data from connection:%s", msg)

		msg = fmt.Sprintf("[%s] say: %s\n", client.Ip, msg)

		broadcastAll(string(msg), client)
	}
}
