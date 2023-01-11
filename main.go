package main

import (
	"fmt"
	"log"

	"techtest/reader"
	"techtest/rpc"
	"techtest/server"

	"github.com/savioxavier/termlink"
)

const (
	DNS_ADDRESS = "A1mgzGvGthnJuScWUgvnYqB4sDcfTixbE1RYeA9tvpP4N4uQjYQ"
)

func main() {

	// Architecture / execution :
	// request to a dns : flappy
	// Resolve dns : flappy -> address
	// Iterate through all chunks / shards to reconstitute the zip file
	// Define a server

	client := rpc.NewClient()

	var dns string

	fmt.Print("Enter the DNS you want to redirect to this server: ")
	fmt.Scan(&dns)
	log.Printf("Resolving dns for flappy...")

	respDns, err := client.ResolveDns(
		DNS_ADDRESS,
		dns,
		"record")
	if err != nil {
		panic(err)
	}

	log.Printf("Result resolve dns for %s => %s", "flappy", string(respDns.Result[0].CandidateValue))

	respData, err := client.Get(
		string(respDns.Result[0].CandidateValue),
		"massa_web_0",
		"",
	)
	if err != nil {
		panic(err)
	}
	// Check the length to be sure it's relevant
	log.Printf("Result Get for %s => length: %d", "A12jcLSmfe9AsRAVJstBoEJNHgr58EtkG9cUUJwUjpsQ5Pes3zA5", len(respData.Result[0].CandidateValue))

	files, err := reader.Read(respData)
	if err != nil {
		panic(err)
	}

	srv := server.DefineServer(files)
	// Run it !
	fmt.Println(termlink.ColorLink("Server running : ", srv.Addr, "italic green"))
	log.Fatal(srv.ListenAndServe())
}
