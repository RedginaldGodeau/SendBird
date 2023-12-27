package main

import (
	"CommandToServe/Client"
	"CommandToServe/Server"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		println("birdserve [client/server] [adress/port]")
		return
	}

	switch arguments[1] {
	case "client":
		var tcp = Client.TCP{Address: arguments[2]}
		err := tcp.Connect()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "server":
		var tcp = Server.TCP{Port: arguments[2]}
		err := tcp.Start()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
