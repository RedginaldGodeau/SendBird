package Client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type TCP struct {
	Address string
}

func (t TCP) Connect() error {
	c, err := net.Dial("tcp", t.Address)
	if err != nil {
		return err
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Print("$ ")
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('§')
		fmt.Print("[SERVER]: " + message)
		if strings.TrimSpace(string(text)) == "EXIT" {
			fmt.Println("Exited...")
			return nil
		}
	}
}
