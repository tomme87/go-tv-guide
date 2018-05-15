package main

import (
	"fmt"
	"github.com/tomme87/go-tv-guide/internal/pkg/config"
	"github.com/tomme87/go-tv-guide/pkg/xmltv"
	"log"
	"os"
)

func main() {
	config.C.Init()

	m := config.C.MongoDB
	m.Connect()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(f)
	}

	var tv xmltv.Tv

	err = tv.LoadXML(f)
	if err != nil {
		log.Fatal(err)
	}

	err = m.InsertChannels(tv.Channels)
	if err != nil {
		log.Fatal(err)
	}

	err = m.InsertProgrammes(tv.Programmes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
