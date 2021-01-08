package main

import (
	"context"
	"fmt"
	"log"
	"syscall/js"

	"github.com/jeffprestes/ethclient"
	"github.com/jeffprestes/ethclient/common"
)

//Client is the object responsible for the conection with Ethereum
var Client *ethclient.Client

func init() {
	var err error
	Client, err = GetCustomNetworkClient("https://mainnet.infura.io/v3/b5b4d5a4220049919ba4765671bad405")
	if err != nil {
		log.Fatalln("Error connecting to Ethereum. ", err.Error())
		return
	}
}

func main()  {
	// js.Global().Set("getBalance", js.FuncOf(GetBalance))
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("GetBalance started")
		go func() {
			balance, err := Client.BalanceAt(context.Background(), common.HexToAddress("0x263C3Ab7E4832eDF623fBdD66ACee71c028Ff591"), nil)
			if err != nil {
				log.Fatalln("Error getting balance. ", err.Error())
			}
			//js.Global().Get("document").Call("getElementById", "loadMsg").Call("alert", balance.String())
			js.Global().Get("window").Call("mostrar", balance.String())
			cb.Release()
		}()
		return nil
	})
	js.Global().Get("document").Call("getElementById", "meubotao").Call("addEventListener", "click", cb)
	canal := make(chan struct{}, 0)
	fmt.Println("Go WebAssembly inicializado")
	<-canal
}

//GetBalance returns ethereum wallet balance
func GetBalance(this js.Value, args []js.Value) interface{} {
	fmt.Println("GetBalance started")
	balance, err := Client.BalanceAt(context.Background(), common.HexToAddress("0x263C3Ab7E4832eDF623fBdD66ACee71c028Ff591"), nil)
	if err != nil {
		log.Fatalln("Error getting balance. ", err.Error())
		return "0"
	}
	return balance.String()
}