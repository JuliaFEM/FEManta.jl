package japp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Model struct {
	modelname string `json:"modelname"`
	// size      int    `json:"size"`
}

type Data struct {
	models []Model `json:"models"`
}

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

	file, _, err := r.FormFile("File")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	randName := RandStringRunes(10)
	filename := "/tmp/" + randName

	// copy example
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, file)

	juliaCmd := "julia -e \"using JuliaFEM; model = AbaqusReader.abaqus_read_model(\"" + filename + "\"); include(\"abaqus.jl\")\""
	cmd := exec.Command("julia", "-e", juliaCmd)
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

// CurrentModelsHandler starts new JuliaFEM calculation
func CurrentModelsHandler(w http.ResponseWriter, r *http.Request) {

	var files []Model

	root := "/tmp"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		model := Model{modelname: path}
		files = append(files, model)
		return nil
	})
	if err != nil {
		panic(err)
	}

	data := Data{models: files}

	json.NewEncoder(w).Encode(data)

}

// ModelHandler starts new JuliaFEM calculation
func ModelHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		// Upload model to filesystem.
		// Now the whole file is expected to be in payload,
		// but maybe just use channels to stream file to filesystem?
		file, handler, err := r.FormFile("File")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		filename := "/tmp/" + handler.Filename

		// copy example
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		io.Copy(f, file)
		return

	} else if r.Method == http.MethodGet {

		// Read model file for client for downloading
		// Maybe stream file?
		buf := make([]byte, 1024)
		vars := mux.Vars(r)

		name := vars["file"]

		filename := "/tmp/" + name

		// copy example
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		for {
			v, _ := f.Read(buf)

			if v == 0 {
				break
			}
			//in case it is a string file, you could check its content here...
			// fmt.Print(string(buf))
		}
		return

	}

}

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
