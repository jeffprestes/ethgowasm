package main

import (
	"log"

	"github.com/jeffprestes/ethclient"
)

/*
GetCustomNetworkClient connects and return a client to user defined Ethereum network
*/
func GetCustomNetworkClient(URL string) (client *ethclient.Client, err error) {
	err = nil
	client, err = ethclient.Dial(URL)
	if err != nil {
		log.Printf("There was a failure connecting to %s: %+v", URL, err)
		return
	}
	return
}

/*
GetCustomNetworkClientWebsocket connects via websocket and return a client to user defined Ethereum network
*/
func GetCustomNetworkClientWebsocket(URL string) (client *ethclient.Client, err error) {
	err = nil
	client, err = ethclient.Dial(URL)
	if err != nil {
		log.Printf("There was a failure connecting to %s via Websocket: %+v", URL, err)
		return
	}
	return
}
