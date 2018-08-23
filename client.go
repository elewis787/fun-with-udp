package main

import (
	"net"
)

// UDPClient - todo
type UDPClient struct {
}

// Send - todo
func (udpc *UDPClient) Send(msg string, address string) error {
	remoteUDPAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	rConn, err := net.DialUDP("udp", nil, remoteUDPAddr)
	if err != nil {
		return err
	}
	if _, err := rConn.Write([]byte(msg)); err != nil {
		return err
	}
	return nil
}
