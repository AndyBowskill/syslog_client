package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := SetupClient("udp", "192.168.48.10:514")
	defer CloseClient(conn)
	if err == nil {
		SendRaw(conn)
	}
}

func SetupClient(network, address string) (net.Conn, error) {

	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func SendRaw(conn net.Conn) error {

	_, err := fmt.Fprint(conn, "Testing test 1, 2, 3")
	if err != nil {
		return err
	}

	return nil
}

func CloseClient(conn net.Conn) {
	conn.Close()
}
