package main

import "github.com/mapleque/cell"

func main() {
	ds := cell.NewDataService()
	server := cell.NewServer(ds, "../tpl/")
	server.Serve(":9999")
}