package main

import (
	"net"
)

type session struct {
	net.Conn
	isAlive bool
}

func (s *session) Start() {

}
