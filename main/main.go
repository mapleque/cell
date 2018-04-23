package main

import "github.com/mapleque/cell"

func main() {
	ds := cell.NewDataService()
	server := cell.NewServer(ds)
	server.Serve(":9999")
}
