package main

import (
	"log"
	"os"
	"strings"

	"github.com/joshua-anderson/rbd/cmd"
	"github.com/joshua-anderson/rbd/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("command is required. Usage: rbd <cmd>")
	}

	config, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	mapping, err := config.GetMap()
	if err != nil {
		log.Fatal(err)
	}

	worker, err := config.GetWorker(mapping.Worker)
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Rsync(worker.Host, worker.User, worker.Port, mapping.Local, mapping.Remote, true)
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Cmd(worker.Host, worker.User, worker.Port, mapping.Remote, strings.Join(os.Args[1:], " "))
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Rsync(worker.Host, worker.User, worker.Port, mapping.Local, mapping.Remote, false)
	if err != nil {
		log.Fatal(err)
	}
}
