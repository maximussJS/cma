package main

import (
	"cma/packages/server"
)

func main() {
	cma := server.NewCMAServer()

	cma.Start()
}
