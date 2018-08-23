package main

import (
	"context"
	"log"
	"net"
)

// UDPServer - todo
type UDPServer struct {
}

// ListenAndServe - todo
func (udps *UDPServer) ListenAndServe(ctx context.Context, address string) error {

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}
	conn, err := net.ListenPacket("udp", udpAddr.String())
	if err != nil {
		return err
	}
	defer conn.Close()

	go func() {
		if err := udps.handleConnection(ctx, conn); err != nil {
			log.Println(err)
		}
	}()
	<-ctx.Done()
	return nil
}

// handleConnection - read bytes - handle bytes - send packet
func (udps *UDPServer) handleConnection(ctx context.Context, conn net.PacketConn) error {
	for {
		inputBytes := make([]byte, 4096)
		_, addr, err := conn.ReadFrom(inputBytes) // Only blocks if input bytes is not zero
		if err != nil {
			return err
		}
		log.Println(addr)
		log.Println("message :", string(inputBytes))
	}
}
