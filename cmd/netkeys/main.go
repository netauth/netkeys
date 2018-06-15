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
	cfg       *client.NACLConfig
	keyType   = flag.String("type", "SSH", "Type of keys to print")
	entityID  = flag.String("ID", "", "ID to look up")
	serviceID = flag.String("service", "netkeys", "Service ID to send")
)

// loadConfig loads the config in.  It would have been nice to do this
// in init(), but that gets called too late
func loadConfig() {
	if cfg != nil {
		return
	}
	var err error
	cfg, err = client.LoadConfig("")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Config loading error: ", err)
		return
	}
}

func main() {
	flag.Parse()

	// Handle config loading
	loadConfig()
	if cfg == nil {
		os.Exit(1)
	}
	cfg.ServiceID = *serviceID

	// Shut off all the logging
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	// Grab the client, we're ignoring the error here since crypto
	// will certainly fail to initialize
	c, err := client.New(cfg)

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
