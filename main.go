package main

import (
	"fmt"
	"retherer/core"
)

func main() {
	fmt.Println("Hi there, this is retherer")
	bc := core.NewBlockchain()

	defer bc.Db.Close()
	cli := core.CLI{Bc: bc}
	cli.Run()
}
