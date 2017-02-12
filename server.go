package main

import "gabby.network/api/server"

func main() {
	s := server.New()
	s.ListenAndServe(":43960")
}
