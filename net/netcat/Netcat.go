package netcat

import (
	"bufio"
	"fmt"
	"net"
)

// Netcat todo
type Netcat struct {
	AddressAndPort string
	WhenAccept     func(string)
	listener       net.Listener
}

// Run todo
func (s *Netcat) Run() {
	listener, err := net.Listen("tcp4", s.AddressAndPort)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	s.listener = listener

	for {
		conn, err := listener.Accept()
		if err != nil || conn == nil {
			fmt.Println("listener err = ", err)
		} else {
			go s.HandleConn(conn)
		}
	}
}

// Close todo
func (s *Netcat) Close() {
	s.listener.Close()
}

// HandleConn todo
func (s *Netcat) HandleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, _, err := reader.ReadLine()
	fmt.Println("line:" + string(line))
	if err != nil {
		fmt.Println("conn err = ", err)
	} else {
		s.WhenAccept(string(line))
	}

}
