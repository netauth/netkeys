package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/NetAuth/NetAuth/pkg/client"
)

var (
	server   = flag.String("server", "localhost", "NetAuth server")
	port     = flag.Int("port", 8080, "NetAuth server port")
	clientID = flag.String("client", "netkeys", "Client ID")
	keyType  = flag.String("type", "SSH", "Type of keys to print")
	entityID = flag.String("ID", "", "ID to look up")
)

func main() {
	flag.Parse()

	// Shut off all the logging
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "hostname-error"
	}

	// Grab the client, we're ignoring the error here since crypto
	// will certainly fail to initialize
	c, err := client.New(*server, *port, *clientID, hostname)

	// This is only ever done for read, never write, so we feed a
	// null token
	keys, err := c.ModifyEntityKeys("", *entityID, "LIST", *keyType, "")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Print out the keys, no formatting, just the key data
	for _, k := range keys {
		parts := strings.Split(k, ":")
		fmt.Println(strings.Join(parts[1:], " "))
	}
	os.Exit(0)
}
