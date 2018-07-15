package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/NetAuth/NetAuth/pkg/client"
)

var (
	cfg       *client.NACLConfig
	keyType   = flag.String("type", "SSH", "Type of keys to print")
	entityID  = flag.String("ID", "", "ID to look up")
	serviceID = flag.String("service", "netkeys", "Service ID to send")
)

func main() {
	flag.Parse()

	// Shut off all the logging
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	// Grab a client
	c, err := client.New(nil)
	if err != nil {
		os.Exit(1)
	}

	// Set the service ID
	c.SetServiceID(*serviceID)

	// This is only ever done for read, never write, so we feed a
	// null token
	keys, err := c.ModifyEntityKeys("", *entityID, "LIST", *keyType, "")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Print out the keys, no formatting, just the key data
	for _, k := range keys {
		fmt.Println(k)
	}
	os.Exit(0)
}
