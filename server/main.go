package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Message struct {
	Sender string
	Message string
	Time string
}

type NameRequest struct {
	Name string
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

var messages *list.List = list.New()
var connections map[*websocket.Conn]*websocket.Conn = make(map[*websocket.Conn]*websocket.Conn)
var usernames map[string]string = make(map[string]string)

func main() {
	log.Println("http://localhost:3000")

	http.HandleFunc("/join", joinHandler)

	http.HandleFunc("/connect", connectionHandler)

	http.Handle("*", http.FileServer(http.Dir("../client/public")))

	http.ListenAndServe(":3000", nil)
}

func joinHandler(response http.ResponseWriter, requestt *http.Request) {
	if requestt.Method == http.MethodPost {
		if body, error := ioutil.ReadAll(requestt.Body); error != nil {
			response.WriteHeader(400)
			response.Write([]byte("Failed to read request body."))
			return
		} else {
			request := &NameRequest{}

			if error := json.Unmarshal(body, request); error != nil {
				response.WriteHeader(400)
				response.Write([]byte("Failed to parse request. Check format."))
				return
			}

			if strings.Compare(usernames[request.Name], request.Name) == 0 {
				response.WriteHeader(401)
				response.Write([]byte("Name already in use."))
				return
			}

			usernames[request.Name] = request.Name

			response.WriteHeader(200)
		}
	}
}

func connectionHandler(response http.ResponseWriter, request *http.Request) {
	if connection := upgrade(response, request); connection != nil {
		connections[connection] = connection

		defer connection.Close()

		for {
			if messageType, data, error := connection.ReadMessage(); error != nil {
				connection.WriteMessage(messageType, []byte(error.Error()))
			} else {
				message := &Message{}

				if error := json.Unmarshal(data, message); error != nil {
					connection.WriteMessage(messageType, []byte(error.Error()))
				} else {
					messages.PushBack(message)

					for key := range connections {
						connections[key].WriteMessage(messageType, []byte(generateJSON(message)))
					}	
				}
			}		

			//check if client is still connected
		}
	}
}

func generateJSON(message *Message) (string) {
	sender := message.Sender
	messageText := message.Message
	time := message.Time

	return fmt.Sprintf(`{ "sender": "%s", "message": "%s", "time": "%s"}`, sender, messageText, time) 
}

func upgrade(response http.ResponseWriter, request *http.Request) (*websocket.Conn) {
	if conn, error := upgrader.Upgrade(response, request, nil); error != nil {
		response.WriteHeader(426)
		response.Write([]byte(error.Error()))
		return nil
	} else {
		return conn
	}
}
