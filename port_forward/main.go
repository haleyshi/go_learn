package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	target string
	port   int
)

func init() {
	flag.StringVar(&target, "target", "", "the target (<host>:<port>)")
	flag.IntVar(&port, "port", 1337, "the tunnel port")
}

func main() {
	// Parse arguments
	flag.Parse()

	// Monitor CTRL+C and quit
	signals := make(chan os.Signal, 1)
	stop := make(chan bool)
	signal.Notify(signals, os.Interrupt)

	go func() {
		for range signals {
			fmt.Println("\nReceived an interrupt, stopping...")
			stop <- true
		}
	}()

	// Start a server on the giving port
	incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Could not start server on %d: %v", port, err)
	}
	fmt.Printf("Server running on %d\n", port)

	// Listen for incoming connections and accept the 1st one
	client, err := incoming.Accept()
	if err != nil {
		log.Fatal("cound not accept client connection", err)
	}

	defer client.Close()

	fmt.Printf("client '%v' connected!\n", client.RemoteAddr())

	// Connect to the target server
	target, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatal("could not connect to target", err)
	}

	defer target.Close()
	fmt.Printf("connection to server %v established!\n", target.RemoteAddr())

	// Pass the traffic between client and target
	go func() { io.Copy(target, client) }()
	go func() { io.Copy(client, target) }()

	// Wait for CTRL+C
	<-stop
}
