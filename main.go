package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	address := "127.0.0.1:7878"
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		<-c
		cancel()
		wg.Done()
	}()

	server := &UDPServer{}
	wg.Add(1)
	go func() {
		if err := server.ListenAndServe(ctx, address); err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	log.Println("setting up server")
	time.Sleep(time.Second * 5) // random amount of time for server to setup
	log.Println("attempting cleint delivery")

	client1 := &UDPClient{}
	client2 := &UDPClient{}
	client3 := &UDPClient{}

	go client1.Send("1", address)
	go client2.Send("2", address)
	go client3.Send("3", address)

	wg.Wait()
}
