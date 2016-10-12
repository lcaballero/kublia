package main

import (
	"encoding/json"
	"fmt"
	"github.com/lcaballero/kublai/cli"
	"github.com/lcaballero/kublai/queue"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	conf := cli.ParseArgs(os.Args...)
	bin, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bin))
	fmt.Printf("%v %v %v %v %v\n",
		queue.Received,
		queue.Published,
		queue.Completed,
		queue.ReadyForArchive,
		queue.ProcessingError,
	)
}
