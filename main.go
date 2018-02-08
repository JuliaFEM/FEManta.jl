package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/juliafem/manta/japp"
)

func main() {

	server := japp.Start()
	log.Println("http server started on :8000")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	defer server.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmt.Printf("start listening\n")
}
