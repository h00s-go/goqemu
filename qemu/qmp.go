package qemu

import (
	"bufio"
	"fmt"
	"net"
)

// QMP is used for communication with QEMU guest
type QMP struct {
	address string
	port    string
}

// NewQMP return new QMP
func NewQMP(address string, port string) *QMP {
	return &QMP{address, port}
}

//SendCommand is sending command to guest using QMP
func (q *QMP) SendCommand(command string) (string, error) {
	conn, err := net.Dial("tcp", q.address+":"+q.port)
	if err != nil {
		return "", err
	}
	fmt.Fprintf(conn, "{ \"execute\": \"qmp_capabilities\" }r\n")
	bufio.NewReader(conn).ReadString('\n')
	fmt.Fprintf(conn, "{ \"execute\": \""+command+"\" }r\n")
	output, err := bufio.NewReader(conn).ReadString('\n')
	return output, err
}
