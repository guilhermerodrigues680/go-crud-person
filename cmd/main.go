package main

import "gocrudperson/internal/api"

const apiAddr string = ":8080"

func main() {
	api.Listen(apiAddr)
}
