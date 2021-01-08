package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main()  {
	http.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request) {
		if request.RequestURI == "/" {
			request.RequestURI = "/index.html"
		}
		writer.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(request.URL.Path, ".wasm") {
			writer.Header().Set("content-type", "application/wasm")
		}
		file, err := os.Open(filepath.Join("./html/", request.RequestURI))
		if err != nil {
			log.Println("Error opening ", request.RequestURI)			
		}
		io.Copy(writer, file)
	})
	port := ":" + os.Getenv("PORT")	
	if port == ":" {
		port = ":4000"
	}
	serverName, err := net.LookupCNAME("ethgowasm.herokuapp.com")
	if err != nil {
		log.Println("Could not get serverName. Error: ", err.Error())
	}
	log.Println("Serving on ", serverName, " on port ", port)
	log.Println(http.ListenAndServe(port, nil))
}