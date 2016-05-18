package namerakad

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Client struct {
	host string
	port int
}

func NewClient() *Client {
	return &Client{host: "localhost", port: 7365}
}

func (c *Client) Run() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", c.host, c.port))
	if err != nil {
		log.Fatalf("failed to connect to %v:%v", c.host, c.port)
	}

	b := bufio.NewReader(conn)
	i := bufio.NewReader(os.Stdin)

	for {
		input, err := i.ReadBytes('\n')
		if err != nil {
			break
		}

		_, err = conn.Write(input)
		if err != nil {
			break
		}

		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}

		log.Print(string(line))
	}
}
