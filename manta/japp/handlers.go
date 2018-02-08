package japp

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var stdoutMessages = make(chan []byte)
var stderrMessages = make(chan []byte)

// ExecuteHandler starts new JuliaFEM calculation
func ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	var errStdout, errStderr error

	// conn, _ := upgrader.Upgrade(w, r, nil)
	fmt.Printf("This is new")

	cmd := exec.Command("julia", "-e", "using JuliaFEM;")
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := PipeFactory(stdoutMessages)
	stderr := PipeFactory(stderrMessages)

	go func() {
		err := cmd.Start()
		defer cmd.Wait()
		if err != nil {
			log.Fatalf("cmd.Start() failed with '%s'\n", err)
		}
	}()

	// Magic happends in the io.Copy function. This triggers Write function
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

}

// IndexHandler starts new JuliaFEM calculation
// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, r.URL.Path)
// }

// WebsocketHandler starts new JuliaFEM calculation
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	for {
		select {
		case stdoutMsg := <-stdoutMessages:
			if err := ws.WriteMessage(websocket.TextMessage, stdoutMsg); err != nil {
				ws.Close()
				break
			}
		case stderrMsg := <-stderrMessages:
			if err := ws.WriteMessage(websocket.TextMessage, stderrMsg); err != nil {
				ws.Close()
				break
			}
		}
	}
}
