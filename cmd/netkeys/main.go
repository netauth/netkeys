package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/spf13/pflag"

	"github.com/NetAuth/NetAuth/pkg/client"
)

var (
	keyType   = pflag.String("type", "SSH", "Type of keys to print")
	entityID  = pflag.String("ID", "", "ID to look up")
	serviceID = pflag.String("service", "netkeys", "Service ID to send")
	cfgfile = pflag.String("config", "", "Config file to use")
	verbose = pflag.Bool("verbose", false, "Show logs")
)

func main() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	if *cfgfile != "" {
		viper.SetConfigFile(*cfgfile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/netauth/")
		viper.AddConfigPath("$HOME/.netauth")
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}
	viper.Set("client.ServiceName", "netauth")

	// Shut off all the logging
	if !*verbose {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	// Grab a client
	c, err := client.New()
	if err != nil {
		fmt.Println("Client initialization error:", err)
		os.Exit(1)
	}

	// Set the service ID
	viper.Set("client.ServiceName", *serviceID)

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
