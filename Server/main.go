package Server

import (
	"CommandToServe/Server/commands"
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type TCP struct {
	Port string
}

func (t TCP) Start() error {
	l, err := net.Listen("tcp", ":"+t.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		return err
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		cmd := string(netData)

		if err != nil {
			return err
		}
		if strings.TrimSpace(cmd) == "exit" {
			fmt.Println("Exiting TCP server!")
			return nil
		}

		switch strings.TrimSpace(cmd) {
		case "status":
			result, err := commands.ContainerStatus()
			printOutput(c, result, err)
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "§"
		c.Write([]byte(myTime))
	}
}

func printOutput(c net.Conn, result string, err error) {
	fmt.Println(result)
	if err != nil {
		c.Write([]byte(err.Error() + "§"))
	} else {
		c.Write([]byte(result + "§"))
	}
}
