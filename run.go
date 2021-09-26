package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"
)

const exitCodeError     = 1
const exitCodeInterrupt = 2

var (
	// check for the --debug flag when running the executable
	DebugMode = flag.Bool("debug", false, "show verbose activity and warning messages")
	// the global app context for handling errors and interrupts
	AppContext context.Context
	// the loaded config.yaml file
	Config *AppConfig
	// the current connected bot session
	Session *discordgo.Session
)

func mainn() {
	// create and store our main app context
	AppContext = context.Background()
	// read and parse the config file
	Config = GetConfig()
	// send a login request to the Discord API
	discord, err := discordgo.New("Bot " + Config.Token)
	if err != nil {
		log.Fatalf("error logging into Discord: %s\n",err)
	}
	log.Printf("logged into Discord!")

	discord.Client.Get("")

	// create an app-wide interrupt signal handler to gracefully shut down
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// do other setup stuff here
	// start the main app loop in a goroutine
	go appLoop(interrupt)
}

func appLoop(signal chan os.Signal) {
	appStartTime := time.Now().Unix()
	(func()  {
		<-signal
		
	})()
	<-signal  //todo: get the reason for exiting
	appEndTime := time.Now().Unix()
	log.Printf("caught interrupt: ran for \n")
}

func appShutdown() {
	log.Printf("shutting down..")
}