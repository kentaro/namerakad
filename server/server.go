package namerakad

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	port int
}

func NewServer() *Server {
	return &Server{port: 7365}
}

func (s *Server) Run() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		log.Panicf("failed to listen on %v", s.port)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("an error occurred while accepting")
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	b := bufio.NewReader(conn)

	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}
		conn.Write(line)
	}
}
